package handlers

import (
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/gorest"
	"github.com/sirupsen/logrus"
)

// MakeListBucketHandler creates handler for listing buckets
func MakeListBucketHandler(db *bolt.DB, logger logrus.StdLogger) gorest.Handler {
	return &gorest.BaseHandler{
		Name: "buckets",
		Path: "/buckets",
		Get: func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			var buckets []string
			var response Response
			defer r.Body.Close()
			err := db.View(func(tx *bolt.Tx) error {
				err := tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
					buckets = append(buckets, string(name))
					return nil
				})
				if nil != err {
					return err
				}
				return nil
			})
			if nil != err {
				logger.Println(err)
				response.AddError(err)
				w.WriteHeader(http.StatusInternalServerError)
				writeResponse(w, response)
				return
			}
			response.Meta.Success = true
			response.Data = buckets
			writeResponse(w, response)
		},
	}
}
