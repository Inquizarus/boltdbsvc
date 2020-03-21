package handlers

import (
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/gorest"
	"github.com/sirupsen/logrus"
)

// MakeBucketHandler creates handler for bucket interactions
func MakeBucketHandler(db *bolt.DB, logger logrus.StdLogger) gorest.Handler {
	return &gorest.BaseHandler{
		Name: "bucket",
		Path: "/buckets/{name}",
		Get: func(_ http.ResponseWriter, r *http.Request, _ map[string]string) {
			defer r.Body.Close()
		},
		Post: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			var response Response
			defer r.Body.Close()
			name := p["name"]
			logger.Printf(`trying to create bucket with name %s`, name)
			err := db.Update(func(tx *bolt.Tx) error {
				_, err := tx.CreateBucket([]byte(name))
				if nil != err {
					return fmt.Errorf("create bucket: %s", err)
				}
				return nil
			})
			if nil != err {
				logger.Println(err)
				response.AddError(err)
				writeResponse(w, response)
				w.WriteHeader(http.StatusConflict)
				return
			}
			logger.Printf(`successfully created bucket with name %s`, name)
			response.Meta.Success = true
			writeResponse(w, response)
		},
		Delete: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			var response Response
			defer r.Body.Close()
			name := p["name"]
			err := db.Update(func(tx *bolt.Tx) error {
				err := tx.DeleteBucket([]byte(name))
				if nil != err {
					return fmt.Errorf("delete bucket: %s", err)
				}
				return nil
			})
			if nil != err {
				logger.Println(err)
				response.AddError(err)
				writeResponse(w, response)
				w.WriteHeader(http.StatusNotFound)
				return
			}
			response.Meta.Success = true
			writeResponse(w, response)
		},
	}
}
