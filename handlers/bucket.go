package handlers

import (
	"net/http"

	"github.com/inquizarus/golbag/storages"
	"github.com/inquizarus/gorest"
	"github.com/sirupsen/logrus"
)

// MakeBucketHandler creates handler for bucket interactions
func MakeBucketHandler(s storages.Storage, logger logrus.StdLogger) gorest.Handler {
	return &gorest.BaseHandler{
		Name: "bucket",
		Path: "/buckets/{name}",
		Get: func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			var response Response
			defer r.Body.Close()
			response.Meta.Message = "not implemented"
			response.Meta.Success = true
			writeResponse(w, response)
		},
		Post: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			var response Response
			defer r.Body.Close()
			name := p["name"]
			logger.Printf(`trying to create bucket with name %s`, name)
			err := s.CreateBucket([]byte(name))
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
			err := s.DeleteBucket([]byte(name))
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
