package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/boltdbsvc/models"
	"github.com/inquizarus/gorest"
	"github.com/sirupsen/logrus"
)

// MakeItemHandler creates handler for item interactions
func MakeItemHandler(db *bolt.DB, logger logrus.StdLogger) gorest.Handler {
	return &gorest.BaseHandler{
		Name: "bucket",
		Path: "/buckets/{bucket_name}/items/{item_name}",
		Get: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			var response Response
			var item models.Item
			defer r.Body.Close()
			bucketName := p["bucket_name"]
			itemName := p["item_name"]
			logger.Println(fmt.Sprintf("trying to retrieve %s from bucket %s", itemName, bucketName))
			err := db.View(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte(bucketName))
				if nil == bucket {
					return fmt.Errorf("bucket with name %s does not exist", bucketName)
				}
				itemBytes := bucket.Get([]byte(itemName))
				return json.Unmarshal(itemBytes, &item)
			})
			if nil == err && 1 > len(item.Content) {
				err = fmt.Errorf("item with name %s does not exist", itemName)
			}
			if nil != err {
				logger.Println(fmt.Errorf("could not retrieve %s from bucket %s: %v", itemName, bucketName, err))
				response.AddError(err)
				w.WriteHeader(http.StatusConflict)
				writeResponse(w, response)
				return
			}
			logger.Println(fmt.Sprintf("successfully retrieved %s from bucket %s", itemName, bucketName))
			w.Write(item.Content)
		},
		Post: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			var response Response
			var item models.Item
			defer r.Body.Close()
			bucketName := p["bucket_name"]
			itemName := p["item_name"]
			body, err := ioutil.ReadAll(r.Body)
			if nil == err {
				logger.Println(fmt.Sprintf("trying to create item %s in bucket %s", itemName, bucketName))
				item.Content = body
				item.Meta.CreatedAt = time.Now().Unix()
				if itemBytes, err := json.Marshal(item); nil == err {
					err = db.Update(func(tx *bolt.Tx) error {
						bucket := tx.Bucket([]byte(bucketName))
						if nil == bucket {
							return fmt.Errorf("bucket with name %s does not exist", bucketName)
						}
						return bucket.Put([]byte(itemName), itemBytes)
					})
				}

			}
			if nil != err {
				logger.Println(fmt.Errorf("could not create item %s in bucket %s: %v", itemName, bucketName, err))
				response.AddError(err)
				w.WriteHeader(http.StatusConflict)
				writeResponse(w, response)
				return
			}
			logger.Println(fmt.Sprintf("successfully created item %s in bucket %s", itemName, bucketName))
			response.Meta.Success = true
			writeResponse(w, response)
		},
		Delete: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			var response Response
			defer r.Body.Close()
			bucketName := p["bucket_name"]
			itemName := p["item_name"]
			logger.Println(fmt.Sprintf("trying to delete %s from bucket %s", itemName, bucketName))
			err := db.Update(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte(bucketName))
				if nil == bucket {
					return fmt.Errorf("bucket with name %s does not exist", bucketName)
				}
				return bucket.Delete([]byte(itemName))
			})
			if nil != err {
				logger.Println(fmt.Errorf("could not delete %s from bucket %s: %v", itemName, bucketName, err))
				response.AddError(err)
				w.WriteHeader(http.StatusInternalServerError)
				writeResponse(w, response)
				return
			}
			logger.Println(fmt.Sprintf("successfully delete %s from bucket %s", itemName, bucketName))
			response.Meta.Success = true
			writeResponse(w, response)
		},
	}
}
