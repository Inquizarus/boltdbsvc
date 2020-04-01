package storages

import (
	"fmt"

	"github.com/inquizarus/golbag/models"
)

// MapStorage handles are storage in memory
type MapStorage struct {
	Map map[string]map[string][]byte
}

// CreateBucket wrapper implementation for map
func (ms *MapStorage) CreateBucket(name []byte) error {
	ms.Map[string(name)] = map[string][]byte{}
	return nil
}

// GetBucket implementation for map
func (ms *MapStorage) GetBucket(name []byte) (models.Bucket, error) {
	var bucket models.Bucket
	if _, ok := ms.Map[string(name)]; ok != true {
		return bucket, fmt.Errorf("could not get bucket")
	}
	return bucket, nil
}

// GetBuckets implementation for map
func (ms *MapStorage) GetBuckets() [][]byte {
	var buckets [][]byte
	for bucket := range ms.Map {
		buckets = append(buckets, []byte(bucket))
	}
	return buckets
}

// DeleteBucket implementation for map
func (ms *MapStorage) DeleteBucket(name []byte) error {
	_, ok := ms.Map[string(name)]
	if ok {
		delete(ms.Map, string(name))
		return nil
	}
	return fmt.Errorf("could not delete bucket")
}

// AddItemToBucket implementation for map
func (ms *MapStorage) AddItemToBucket(k []byte, b []byte, i []byte) error {
	_, err := ms.GetBucket(b)
	if nil != err {
		return fmt.Errorf("AddItemToBucket: %v", err)
	}
	ms.Map[string(b)][string(k)] = i
	return nil
}

// GetItemFromBucket implementation for map
func (ms *MapStorage) GetItemFromBucket(k []byte, b []byte) ([]byte, error) {
	_, err := ms.GetBucket(b)
	if nil != err {
		return []byte{}, fmt.Errorf("GetItemFromBucket: %v", err)
	}
	ib, ok := ms.Map[string(b)][string(k)]
	if true != ok {
		return []byte{}, fmt.Errorf("Could not get item with key %s from bucket %s", k, b)
	}
	return ib, nil
}

// DeleteItemFromBucket implementation for map
func (ms *MapStorage) DeleteItemFromBucket(k []byte, b []byte) error {
	_, err := ms.GetBucket(b)
	if nil != err {
		return fmt.Errorf("DeleteItemFromBucket: %v", err)
	}
	if _, ok := ms.Map[string(b)][string(k)]; ok {
		delete(ms.Map[string(b)], string(k))
		return nil
	}
	return fmt.Errorf("Could not get item with key %s from bucket %s", k, b)
}
