// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/s3"
)

type FakeUnversionedBucket struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	RegionNameStub        func() string
	regionNameMutex       sync.RWMutex
	regionNameArgsForCall []struct{}
	regionNameReturns     struct {
		result1 string
	}
	regionNameReturnsOnCall map[int]struct {
		result1 string
	}
	CopyObjectStub        func(key, originPath, destinationPath, originBucketName, originBucketRegion string) error
	copyObjectMutex       sync.RWMutex
	copyObjectArgsForCall []struct {
		key                string
		originPath         string
		destinationPath    string
		originBucketName   string
		originBucketRegion string
	}
	copyObjectReturns struct {
		result1 error
	}
	copyObjectReturnsOnCall map[int]struct {
		result1 error
	}
	ListFilesStub        func(path string) ([]string, error)
	listFilesMutex       sync.RWMutex
	listFilesArgsForCall []struct {
		path string
	}
	listFilesReturns struct {
		result1 []string
		result2 error
	}
	listFilesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnversionedBucket) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeUnversionedBucket) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeUnversionedBucket) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnversionedBucket) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnversionedBucket) RegionName() string {
	fake.regionNameMutex.Lock()
	ret, specificReturn := fake.regionNameReturnsOnCall[len(fake.regionNameArgsForCall)]
	fake.regionNameArgsForCall = append(fake.regionNameArgsForCall, struct{}{})
	fake.recordInvocation("RegionName", []interface{}{})
	fake.regionNameMutex.Unlock()
	if fake.RegionNameStub != nil {
		return fake.RegionNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.regionNameReturns.result1
}

func (fake *FakeUnversionedBucket) RegionNameCallCount() int {
	fake.regionNameMutex.RLock()
	defer fake.regionNameMutex.RUnlock()
	return len(fake.regionNameArgsForCall)
}

func (fake *FakeUnversionedBucket) RegionNameReturns(result1 string) {
	fake.RegionNameStub = nil
	fake.regionNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnversionedBucket) RegionNameReturnsOnCall(i int, result1 string) {
	fake.RegionNameStub = nil
	if fake.regionNameReturnsOnCall == nil {
		fake.regionNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.regionNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnversionedBucket) CopyObject(key string, originPath string, destinationPath string, originBucketName string, originBucketRegion string) error {
	fake.copyObjectMutex.Lock()
	ret, specificReturn := fake.copyObjectReturnsOnCall[len(fake.copyObjectArgsForCall)]
	fake.copyObjectArgsForCall = append(fake.copyObjectArgsForCall, struct {
		key                string
		originPath         string
		destinationPath    string
		originBucketName   string
		originBucketRegion string
	}{key, originPath, destinationPath, originBucketName, originBucketRegion})
	fake.recordInvocation("CopyObject", []interface{}{key, originPath, destinationPath, originBucketName, originBucketRegion})
	fake.copyObjectMutex.Unlock()
	if fake.CopyObjectStub != nil {
		return fake.CopyObjectStub(key, originPath, destinationPath, originBucketName, originBucketRegion)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyObjectReturns.result1
}

func (fake *FakeUnversionedBucket) CopyObjectCallCount() int {
	fake.copyObjectMutex.RLock()
	defer fake.copyObjectMutex.RUnlock()
	return len(fake.copyObjectArgsForCall)
}

func (fake *FakeUnversionedBucket) CopyObjectArgsForCall(i int) (string, string, string, string, string) {
	fake.copyObjectMutex.RLock()
	defer fake.copyObjectMutex.RUnlock()
	return fake.copyObjectArgsForCall[i].key, fake.copyObjectArgsForCall[i].originPath, fake.copyObjectArgsForCall[i].destinationPath, fake.copyObjectArgsForCall[i].originBucketName, fake.copyObjectArgsForCall[i].originBucketRegion
}

func (fake *FakeUnversionedBucket) CopyObjectReturns(result1 error) {
	fake.CopyObjectStub = nil
	fake.copyObjectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUnversionedBucket) CopyObjectReturnsOnCall(i int, result1 error) {
	fake.CopyObjectStub = nil
	if fake.copyObjectReturnsOnCall == nil {
		fake.copyObjectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyObjectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUnversionedBucket) ListFiles(path string) ([]string, error) {
	fake.listFilesMutex.Lock()
	ret, specificReturn := fake.listFilesReturnsOnCall[len(fake.listFilesArgsForCall)]
	fake.listFilesArgsForCall = append(fake.listFilesArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("ListFiles", []interface{}{path})
	fake.listFilesMutex.Unlock()
	if fake.ListFilesStub != nil {
		return fake.ListFilesStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listFilesReturns.result1, fake.listFilesReturns.result2
}

func (fake *FakeUnversionedBucket) ListFilesCallCount() int {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	return len(fake.listFilesArgsForCall)
}

func (fake *FakeUnversionedBucket) ListFilesArgsForCall(i int) string {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	return fake.listFilesArgsForCall[i].path
}

func (fake *FakeUnversionedBucket) ListFilesReturns(result1 []string, result2 error) {
	fake.ListFilesStub = nil
	fake.listFilesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeUnversionedBucket) ListFilesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ListFilesStub = nil
	if fake.listFilesReturnsOnCall == nil {
		fake.listFilesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listFilesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeUnversionedBucket) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.regionNameMutex.RLock()
	defer fake.regionNameMutex.RUnlock()
	fake.copyObjectMutex.RLock()
	defer fake.copyObjectMutex.RUnlock()
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnversionedBucket) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ s3.UnversionedBucket = new(FakeUnversionedBucket)
