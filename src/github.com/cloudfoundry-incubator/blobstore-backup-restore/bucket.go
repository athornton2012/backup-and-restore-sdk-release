package blobstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//go:generate counterfeiter -o fakes/fake_bucket.go . Bucket
type Bucket interface {
	Name() string
	RegionName() string
	Versions() ([]Version, error)
	PutVersions(regionName, bucketName string, versions []LatestVersion) error
}

type S3Bucket struct {
	awsCliPath string
	name       string
	regionName string
	accessKey  S3AccessKey
	endpoint   string
}

type S3AccessKey struct {
	Id     string
	Secret string
}

func NewS3Bucket(awsCliPath, name, region, endpoint string, accessKey S3AccessKey) S3Bucket {
	return S3Bucket{
		awsCliPath: awsCliPath,
		name:       name,
		regionName: region,
		accessKey:  accessKey,
		endpoint:   endpoint,
	}
}

func (b S3Bucket) Name() string {
	return b.name
}

func (b S3Bucket) RegionName() string {
	return b.regionName
}

func (b S3Bucket) Versions() ([]Version, error) {
	s3Cli := NewS3CLI(b.awsCliPath, b.endpoint, b.regionName, b.accessKey.Id, b.accessKey.Secret)
	output, err := s3Cli.ListObjectVersions(b.name)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(string(output)) == "" {
		return []Version{}, nil
	}

	response := S3ListVersionsResponse{}
	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}

	return response.Versions, nil
}

func (b S3Bucket) PutVersions(regionName, bucketName string, versions []LatestVersion) error {
	var err error

	awsSession, err := session.NewSession(&aws.Config{
		Region:      &regionName,
		Credentials: credentials.NewStaticCredentials(b.accessKey.Id, b.accessKey.Secret, ""),
		Endpoint:    &b.endpoint,
	})

	if err != nil {
		return err
	}

	s3Session := s3.New(awsSession)

	for _, version := range versions {
		err = b.putVersion(s3Session, regionName, bucketName, version)
		if err != nil {
			return err
		}
	}

	files, err := b.listFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		included := versionsIncludeFile(file, versions)
		if !included {
			err = b.deleteObject(s3Session, b.regionName, b.name, file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b S3Bucket) deleteObject(s3Session *s3.S3, regionName, bucketName, file string) error {
	input := s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file),
	}

	_, err := s3Session.DeleteObject(&input)

	return err
}

func (b S3Bucket) putVersion(s3Session *s3.S3, regionName, bucketName string, version LatestVersion) error {
	input := s3.CopyObjectInput{
		Bucket:     aws.String(bucketName),
		Key:        aws.String(version.BlobKey),
		CopySource: aws.String(fmt.Sprintf("/%s/%s?versionId=%s", bucketName, version.BlobKey, version.Id)),
	}

	_, err := s3Session.CopyObject(&input)

	return err
}

func (b S3Bucket) listFiles() ([]string, error) {
	b
	output, err := b.runS3CLICommand("list-objects", "--bucket", b.name)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(string(output)) == "" {
		return []string{}, nil
	}

	response := S3ListResponse{}
	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}

	files := []string{}
	for _, object := range response.Contents {
		files = append(files, object.Key)
	}

	return files, nil
}

func (b S3Bucket) runS3CLICommand(args ...string) ([]byte, error) {
	outputBuffer := new(bytes.Buffer)
	errorBuffer := new(bytes.Buffer)

	var baseArgs []string
	if b.endpoint != "" {
		baseArgs = []string{"--output", "json", "--region", b.regionName, "--endpoint", b.endpoint, "s3api"}
	} else {
		baseArgs = []string{"--output", "json", "--region", b.regionName, "s3api"}
	}

	awsCmd := exec.Command(b.awsCliPath, append(baseArgs, args...)...)
	awsCmd.Env = append(awsCmd.Env, "AWS_ACCESS_KEY_ID="+b.accessKey.Id)
	awsCmd.Env = append(awsCmd.Env, "AWS_SECRET_ACCESS_KEY="+b.accessKey.Secret)
	awsCmd.Stdout = outputBuffer
	awsCmd.Stderr = errorBuffer

	err := awsCmd.Run()
	if err != nil {
		return nil, errors.New(errorBuffer.String())
	}

	return outputBuffer.Bytes(), nil
}

type S3ListVersionsResponse struct {
	Versions []Version
}

type Version struct {
	Key      string
	Id       string `json:"VersionId"`
	IsLatest bool
}

type S3ListResponse struct {
	Contents []Object
}

type Object struct {
	Key string
}

func versionsIncludeFile(file string, versions []LatestVersion) bool {
	for _, version := range versions {
		if version.BlobKey == file {
			return true
		}
	}

	return false
}

type S3Cli struct {
	awsCliPath      string
	endpoint        string
	regionName      string
	accessKeyId     string
	accessKeySecret string
}

func (s3Cli S3Cli) ListObjects(s string) ([]byte, error) {
	return s3Cli.run("list-objects", "--bucket", s)
}

func (s3Cli S3Cli) ListObjectVersions(s string) ([]byte, error) {
	return s3Cli.run("list-object-versions", "--bucket", s)
}

func (s3Cli S3Cli) run(args ...string) ([]byte, error) {
	outputBuffer := new(bytes.Buffer)
	errorBuffer := new(bytes.Buffer)

	var baseArgs []string
	if s3Cli.endpoint != "" {
		baseArgs = []string{"--output", "json", "--region", s3Cli.regionName, "--endpoint", s3Cli.endpoint, "s3api"}
	} else {
		baseArgs = []string{"--output", "json", "--region", s3Cli.regionName, "s3api"}
	}

	awsCmd := exec.Command(s3Cli.awsCliPath, append(baseArgs, args...)...)
	awsCmd.Env = append(awsCmd.Env, "AWS_ACCESS_KEY_ID="+s3Cli.accessKeyId)
	awsCmd.Env = append(awsCmd.Env, "AWS_SECRET_ACCESS_KEY="+s3Cli.accessKeySecret)
	awsCmd.Stdout = outputBuffer
	awsCmd.Stderr = errorBuffer

	err := awsCmd.Run()
	if err != nil {
		return nil, errors.New(errorBuffer.String())
	}

	return outputBuffer.Bytes(), nil
}

func NewS3CLI(awsCliPath, endpoint, regionName, accessKeyId, accessKeySecret string) S3Cli {
	return S3Cli{
		awsCliPath:      awsCliPath,
		endpoint:        endpoint,
		regionName:      regionName,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
	}
}
