package storages

import "github.com/inquizarus/boltdbsvc/models"

// Storage is the interface for any types of storage when
// dealing with service data
type Storage interface {
	CreateBucket([]byte) error
	GetBucket([]byte) (models.Bucket, error)
	GetBuckets() [][]byte
	DeleteBucket([]byte) error
	AddItemToBucket([]byte, []byte, []byte) error
	GetItemFromBucket([]byte, []byte) ([]byte, error)
	DeleteItemFromBucket([]byte, []byte) error
}
