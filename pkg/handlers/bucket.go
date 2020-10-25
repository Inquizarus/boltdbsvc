package handlers

import (
	"fmt"
	"net/http"

	"github.com/inquizarus/golbag/pkg/logging"
	"github.com/inquizarus/golbag/pkg/storages"
	"github.com/inquizarus/gorest"
)

// MakeBucketHandler creates handler for bucket interactions
func MakeBucketHandler(s storages.Storage, log logging.Logger) gorest.Handler {
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
			log.Debug(fmt.Sprintf(`trying to create bucket with name %s`, name))
			err := s.CreateBucket([]byte(name))
			if nil != err {
				log.Error(err)
				response.AddError(err)
				writeResponse(w, response)
				w.WriteHeader(http.StatusConflict)
				return
			}
			log.Debug(fmt.Sprintf(`successfully created bucket with name %s`, name))
			response.Meta.Success = true
			writeResponse(w, response)
		},
		Delete: func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			var response Response
			defer r.Body.Close()
			name := p["name"]
			err := s.DeleteBucket([]byte(name))
			if nil != err {
				log.Error(err)
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
