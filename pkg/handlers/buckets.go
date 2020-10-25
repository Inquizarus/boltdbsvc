package handlers

import (
	"net/http"

	"github.com/inquizarus/golbag/pkg/logging"
	"github.com/inquizarus/golbag/pkg/storages"
	"github.com/inquizarus/gorest"
)

// MakeListBucketHandler creates handler for listing buckets
func MakeListBucketHandler(s storages.Storage, log logging.Logger) gorest.Handler {
	return &gorest.BaseHandler{
		Name: "buckets",
		Path: "/buckets",
		Get: func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			var response Response
			var stringBuckets []string
			defer r.Body.Close()
			buckets := s.GetBuckets()
			for _, bucket := range buckets {
				stringBuckets = append(stringBuckets, string(bucket))
			}
			response.Meta.Success = true
			response.Data = stringBuckets
			writeResponse(w, response)
		},
	}
}
