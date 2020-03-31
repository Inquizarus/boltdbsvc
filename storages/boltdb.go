package storages

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/boltdbsvc/models"
)

// BoltDBStorage wrapper implementation for boltdb
type BoltDBStorage struct {
	db *bolt.DB
}

// CreateBucket wrapper implementation for boltdb
func (bs *BoltDBStorage) CreateBucket(name []byte) error {
	err := bs.db.Update(func(t *bolt.Tx) error {
		_, bucketErr := t.CreateBucket(name)
		return bucketErr
	})
	if nil != err {
		return fmt.Errorf("CreateBucket: %v", err)
	}
	return err
}

// GetBucket implementation for boltdb
func (bs *BoltDBStorage) GetBucket(name []byte) (models.Bucket, error) {
	err := bs.db.View(func(t *bolt.Tx) error {
		b := t.Bucket(name)
		if nil != b {
			return nil
		}
		return fmt.Errorf("could not get bucket with name %s, does it exist?", name)
	})
	if nil != err {
		err = fmt.Errorf("GetBucket: %v", err)
	}
	return models.Bucket{}, err
}

// DeleteBucket implementation for boltdb
func (bs *BoltDBStorage) DeleteBucket(name []byte) error {
	err := bs.db.Update(func(t *bolt.Tx) error {
		return t.DeleteBucket(name)
	})
	if nil != err {
		return fmt.Errorf("DeleteBucket: %v", err)
	}
	return err
}

// AddItemToBucket implementation for boltdb
func (bs *BoltDBStorage) AddItemToBucket(k []byte, b []byte, i []byte) error {
	_, err := bs.GetBucket(b)
	if nil != err {
		return fmt.Errorf("AddItemToBucket: %v", err)
	}
	err = bs.db.Update(func(t *bolt.Tx) error {
		return t.Bucket(b).Put(k, i)
	})
	if nil != err {
		return fmt.Errorf("AddItemToBucket: %v", err)
	}
	return nil
}

// GetItemFromBucket implementation for boltdb
func (bs *BoltDBStorage) GetItemFromBucket(k []byte, b []byte) ([]byte, error) {
	var v []byte
	var err error
	_, err = bs.GetBucket(b)
	if nil != err {
		return v, fmt.Errorf("GetItemFromBucket: %v", err)
	}
	err = bs.db.View(func(t *bolt.Tx) error {
		v = t.Bucket(b).Get(k)
		if nil == v {
			return fmt.Errorf("could not find any content under the key %s", k)
		}
		return nil
	})
	if nil != err {
		return v, fmt.Errorf("GetItemFromBucket: %v", err)
	}
	return v, err
}

// DeleteItemFromBucket implementation for boltdb
func (bs *BoltDBStorage) DeleteItemFromBucket(k []byte, b []byte) error {
	_, err := bs.GetBucket(b)
	if nil != err {
		return fmt.Errorf("DeleteItemFromBucket: %v", err)
	}
	err = bs.db.Update(func(t *bolt.Tx) error {
		return t.Bucket(b).Delete(k)
	})
	if nil != err {
		return fmt.Errorf("DeleteItemFromBucket: %v", err)
	}
	return nil
}

// MakeBoltDBStorage is a wrapper for injecting db into the Storage
// struct to not allowing it to leak later
func MakeBoltDBStorage(db *bolt.DB) Storage {
	return &BoltDBStorage{
		db: db,
	}
}
