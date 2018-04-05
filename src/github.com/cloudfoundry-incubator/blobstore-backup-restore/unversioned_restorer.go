package blobstore

import "fmt"

type UnversionedRestorer struct {
	bucketPairs map[string]UnversionedBucketPair
	artifact    UnversionedArtifact
}

func NewUnversionedRestorer(bucketPairs map[string]UnversionedBucketPair, artifact UnversionedArtifact) UnversionedRestorer {
	return UnversionedRestorer{
		bucketPairs: bucketPairs,
		artifact:    artifact,
	}
}

func (b UnversionedRestorer) Run() error {
	backupBucketAddresses, err := b.artifact.Load()
	if err != nil {
		return err
	}

	for key, pair := range b.bucketPairs {
		_, exists := backupBucketAddresses[key]
		if !exists {
			return fmt.Errorf("cannot restore bucket %s, not found in backup artifact", key)
		}
		err = pair.Restore(backupBucketAddresses[key].Path)
		if err != nil {
			return err
		}
	}
	return nil
}
