package storages_test

import (
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/boltdbsvc/storages"
	"github.com/stretchr/testify/assert"
)

const (
	boltDbPath = "./testdb"
)

func TestBoltDBItReturnsErrorWhenBucketDoesNotExist(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	_, err := s.GetBucket([]byte(bucket))
	assert.NotNil(t, err)
	os.Remove(boltDbPath)
}

func TestBoltDBItCanGetBucket(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	s.CreateBucket([]byte(bucket))
	_, err := s.GetBucket([]byte(bucket))
	assert.Nil(t, err)
	os.Remove(boltDbPath)
}

func TestBoltDBItCanGetBuckets(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	s.CreateBucket([]byte(bucket))
	buckets := s.GetBuckets()
	assert.NotEmpty(t, buckets)
	os.Remove(boltDbPath)
}

func TestBoltDBItCanCreateBucket(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	err := s.CreateBucket([]byte(bucket))
	assert.Nil(t, err)
	os.Remove(boltDbPath)
}

func TestBoltDBThatItCanDeleteBucket(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	_ = s.CreateBucket([]byte(bucket))
	assert.Nil(t, s.DeleteBucket([]byte(bucket)))
	os.Remove(boltDbPath)
}

func TestBoltDBThatItReturnsErrorWhenBucketDoesNotExist(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	assert.NotNil(t, s.DeleteBucket([]byte(bucket)))
	os.Remove(boltDbPath)
}

func TestBoltDBThatItCanAddItemToBucket(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	_ = s.CreateBucket([]byte(bucket))
	err := s.AddItemToBucket([]byte(key), []byte(bucket), []byte(value))
	assert.Nil(t, err)
	os.Remove(boltDbPath)
}

func TestBoltDBThatErrorIsReturnedWhenBucketIsNotDefined(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	err := s.AddItemToBucket([]byte(key), []byte(bucket), []byte(value))
	assert.NotNil(t, err)
	os.Remove(boltDbPath)
}

func TestBoltDBThatItCanGetItemFromBucket(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	s.CreateBucket([]byte(bucket))
	s.AddItemToBucket([]byte(key), []byte(bucket), []byte(value))
	ib, err := s.GetItemFromBucket([]byte(key), []byte(bucket))
	assert.Nil(t, err)
	assert.Equal(t, []byte(value), ib)
	os.Remove(boltDbPath)
}

func TestBoltDBGetItemErrorWhenBucketDontExist(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	ib, err := s.GetItemFromBucket([]byte(key), []byte(bucket))
	assert.NotNil(t, err)
	assert.Empty(t, ib)
	os.Remove(boltDbPath)
}

func TestBoltDBGetItemErrorWhenItemDontExist(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	s.CreateBucket([]byte(bucket))
	ib, err := s.GetItemFromBucket([]byte(key), []byte(bucket))
	assert.NotNil(t, err)
	assert.Empty(t, ib)
	os.Remove(boltDbPath)
}

func TestBoltDBDeleteItemReturnNilOnSuccess(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	s.CreateBucket([]byte(bucket))
	s.AddItemToBucket([]byte(key), []byte(bucket), []byte(value))
	err := s.DeleteItemFromBucket([]byte(key), []byte(bucket))
	assert.Nil(t, err)
	os.Remove(boltDbPath)
}

func TestBoltDBDeleteItemErrorWhenBucketDontExist(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	err := s.DeleteItemFromBucket([]byte(key), []byte(bucket))
	assert.NotNil(t, err)
	os.Remove(boltDbPath)
}

func TestBoltDBDeleteItemNilWhenItemDontExist(t *testing.T) {
	db, _ := bolt.Open(boltDbPath, 0600, nil)
	s := storages.MakeBoltDBStorage(db)
	s.CreateBucket([]byte(bucket))
	err := s.DeleteItemFromBucket([]byte(key), []byte(bucket))
	assert.Nil(t, err)
	os.Remove(boltDbPath)
}
