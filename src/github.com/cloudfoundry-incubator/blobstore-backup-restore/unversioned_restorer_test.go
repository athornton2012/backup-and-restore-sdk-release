package blobstore_test

import (
	"fmt"

	. "github.com/cloudfoundry-incubator/blobstore-backup-restore"
	"github.com/cloudfoundry-incubator/blobstore-backup-restore/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("UnversionedRestorer", func() {
	var (
		dropletsBucketPair *fakes.FakeUnversionedBucketPair
		packagesBucketPair *fakes.FakeUnversionedBucketPair

		bucketPairs map[string]UnversionedBucketPair

		artifact *fakes.FakeUnversionedArtifact

		err error

		restorer UnversionedRestorer
	)

	BeforeEach(func() {
		dropletsBucketPair = new(fakes.FakeUnversionedBucketPair)
		packagesBucketPair = new(fakes.FakeUnversionedBucketPair)

		artifact = new(fakes.FakeUnversionedArtifact)
		artifact.LoadReturns(map[string]BackupBucketAddress{
			"droplets": {
				BucketName:   "artifact_backup_droplet_bucket",
				BucketRegion: "artifact_backup_droplet_region",
				Path:         "timestamp/droplets",
			},
			"packages": {
				BucketName:   "artifact_backup_package_bucket",
				BucketRegion: "artifact_backup_package_region",
				Path:         "timestamp2/packages",
			},
		}, nil)

		bucketPairs = map[string]UnversionedBucketPair{
			"droplets": dropletsBucketPair,
			"packages": packagesBucketPair,
		}

		restorer = NewUnversionedRestorer(bucketPairs, artifact)
	})

	JustBeforeEach(func() {
		err = restorer.Run()
	})

	Context("When the artifact is valid and copying works", func() {
		BeforeEach(func() {
			dropletsBucketPair.RestoreReturns(nil)
			packagesBucketPair.RestoreReturns(nil)

		})

		It("restores all the bucket pairs", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(dropletsBucketPair.RestoreCallCount()).To(Equal(1))
			Expect(dropletsBucketPair.RestoreArgsForCall(0)).To(Equal("timestamp/droplets"))
			Expect(packagesBucketPair.RestoreCallCount()).To(Equal(1))
			Expect(packagesBucketPair.RestoreArgsForCall(0)).To(Equal("timestamp2/packages"))
		})
	})

	Context("When a bucket cannot be restored", func() {
		BeforeEach(func() {
			dropletsBucketPair.RestoreReturns(fmt.Errorf("restore error"))
		})

		It("returns an error", func() {
			Expect(err).To(MatchError("restore error"))
		})
	})

	Context("when the artifact cannot be loaded", func() {
		BeforeEach(func() {
			artifact.LoadReturns(nil, fmt.Errorf("CANNOT LOAD ARTIFACT"))
		})

		It("returns an error", func() {
			Expect(err).To(MatchError("CANNOT LOAD ARTIFACT"))
		})
	})

	Context("When there is a mismatch between the backup artifact and the bucket pairs", func() {
		var notInArtifactPair *fakes.FakeUnversionedBucketPair

		BeforeEach(func() {
			notInArtifactPair = new(fakes.FakeUnversionedBucketPair)

			bucketPairs = map[string]UnversionedBucketPair{
				"droplets":        dropletsBucketPair,
				"packages":        packagesBucketPair,
				"not-in-artifact": notInArtifactPair,
			}
		})

		It("returns an error", func() {
			Expect(err).To(MatchError("cannot restore bucket not-in-artifact, not found in backup artifact"))
		})
	})
})
