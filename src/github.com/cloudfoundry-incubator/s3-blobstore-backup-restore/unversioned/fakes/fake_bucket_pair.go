// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/unversioned"
)

type FakeBucketPair struct {
	BackupStub        func(backupLocation string) (unversioned.BackupBucketAddress, error)
	backupMutex       sync.RWMutex
	backupArgsForCall []struct {
		backupLocation string
	}
	backupReturns struct {
		result1 unversioned.BackupBucketAddress
		result2 error
	}
	backupReturnsOnCall map[int]struct {
		result1 unversioned.BackupBucketAddress
		result2 error
	}
	RestoreStub        func(backupLocation string) error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct {
		backupLocation string
	}
	restoreReturns struct {
		result1 error
	}
	restoreReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBucketPair) Backup(backupLocation string) (unversioned.BackupBucketAddress, error) {
	fake.backupMutex.Lock()
	ret, specificReturn := fake.backupReturnsOnCall[len(fake.backupArgsForCall)]
	fake.backupArgsForCall = append(fake.backupArgsForCall, struct {
		backupLocation string
	}{backupLocation})
	fake.recordInvocation("Backup", []interface{}{backupLocation})
	fake.backupMutex.Unlock()
	if fake.BackupStub != nil {
		return fake.BackupStub(backupLocation)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.backupReturns.result1, fake.backupReturns.result2
}

func (fake *FakeBucketPair) BackupCallCount() int {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return len(fake.backupArgsForCall)
}

func (fake *FakeBucketPair) BackupArgsForCall(i int) string {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return fake.backupArgsForCall[i].backupLocation
}

func (fake *FakeBucketPair) BackupReturns(result1 unversioned.BackupBucketAddress, result2 error) {
	fake.BackupStub = nil
	fake.backupReturns = struct {
		result1 unversioned.BackupBucketAddress
		result2 error
	}{result1, result2}
}

func (fake *FakeBucketPair) BackupReturnsOnCall(i int, result1 unversioned.BackupBucketAddress, result2 error) {
	fake.BackupStub = nil
	if fake.backupReturnsOnCall == nil {
		fake.backupReturnsOnCall = make(map[int]struct {
			result1 unversioned.BackupBucketAddress
			result2 error
		})
	}
	fake.backupReturnsOnCall[i] = struct {
		result1 unversioned.BackupBucketAddress
		result2 error
	}{result1, result2}
}

func (fake *FakeBucketPair) Restore(backupLocation string) error {
	fake.restoreMutex.Lock()
	ret, specificReturn := fake.restoreReturnsOnCall[len(fake.restoreArgsForCall)]
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct {
		backupLocation string
	}{backupLocation})
	fake.recordInvocation("Restore", []interface{}{backupLocation})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub(backupLocation)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.restoreReturns.result1
}

func (fake *FakeBucketPair) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeBucketPair) RestoreArgsForCall(i int) string {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return fake.restoreArgsForCall[i].backupLocation
}

func (fake *FakeBucketPair) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucketPair) RestoreReturnsOnCall(i int, result1 error) {
	fake.RestoreStub = nil
	if fake.restoreReturnsOnCall == nil {
		fake.restoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucketPair) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBucketPair) recordInvocation(key string, args []interface{}) {
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

var _ unversioned.BucketPair = new(FakeBucketPair)
