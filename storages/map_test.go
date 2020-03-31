package storages_test

import (
	"testing"

	"github.com/inquizarus/boltdbsvc/storages"
	"github.com/stretchr/testify/assert"
)

const (
	bucket = "bucket1"
	key    = "key1"
	value  = "value1"
)

func TestItCanGetBucket(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			bucket: map[string][]byte{
				key: []byte(value),
			},
		},
	}
	_, err := s.GetBucket([]byte(bucket))
	assert.Nil(t, err)
}

func TestItReturnsErrorWhenBucketDoesNotExist(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{},
	}
	_, err := s.GetBucket([]byte(bucket))
	assert.NotNil(t, err)
}

func TestItCanCreateBucket(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{},
	}
	s.CreateBucket([]byte(bucket))
	_, err := s.GetBucket([]byte(bucket))
	assert.Nil(t, err)
}

func TestThatItCanDeleteBucket(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			bucket: map[string][]byte{
				key: []byte(value),
			},
		},
	}
	assert.Nil(t, s.DeleteBucket([]byte(bucket)))
}

func TestThatItReturnsErrorWhenBucketDoesNotExist(t *testing.T) {
	s := storages.MapStorage{}
	assert.NotNil(t, s.DeleteBucket([]byte(bucket)))
}

func TestThatItCanAddItemToBucket(t *testing.T) {
	m := map[string]map[string][]byte{
		bucket: map[string][]byte{},
	}
	s := storages.MapStorage{
		Map: m,
	}
	err := s.AddItemToBucket([]byte(key), []byte(bucket), []byte(value))
	assert.Nil(t, err)
	if _, ok := m[bucket][key]; ok != true {
		t.Error("item was not added to bucket correctly")
	}
}

func TestThatErrorIsReturnedWhenBucketIsNotDefined(t *testing.T) {
	m := map[string]map[string][]byte{}
	s := storages.MapStorage{
		Map: m,
	}
	err := s.AddItemToBucket([]byte(key), []byte(bucket), []byte(value))
	assert.NotNil(t, err)
}

func TestThatItCanGetItemFromBucket(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			bucket: map[string][]byte{
				key: []byte(value),
			},
		},
	}
	ib, err := s.GetItemFromBucket([]byte(key), []byte(bucket))
	assert.Nil(t, err)
	assert.Equal(t, []byte(value), ib)
}

func TestGetItemErrorWhenBucketDontExist(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{},
	}
	ib, err := s.GetItemFromBucket([]byte(key), []byte(bucket))
	assert.NotNil(t, err)
	assert.Empty(t, ib)
}

func TestGetItemErrorWhenItemDontExist(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			bucket: map[string][]byte{},
		},
	}
	ib, err := s.GetItemFromBucket([]byte(key), []byte(bucket))
	assert.NotNil(t, err)
	assert.Empty(t, ib)
}

func TestDeleteItemReturnNilOnSuccess(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			bucket: map[string][]byte{
				key: []byte(value),
			},
		},
	}
	err := s.DeleteItemFromBucket([]byte(key), []byte(bucket))
	assert.Nil(t, err)
}

func TestDeleteItemErrorWhenBucketDontExist(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{},
	}
	err := s.DeleteItemFromBucket([]byte(key), []byte(bucket))
	assert.NotNil(t, err)
}

func TestDeleteItemErrorWhenItemDontExist(t *testing.T) {
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			bucket: map[string][]byte{},
		},
	}
	err := s.DeleteItemFromBucket([]byte(key), []byte(bucket))
	assert.NotNil(t, err)
}
