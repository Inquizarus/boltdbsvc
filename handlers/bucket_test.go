package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/inquizarus/golbag/handlers"
	"github.com/inquizarus/golbag/storages"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestThatTheBucketHandlerWorks(t *testing.T) {
	logger := logrus.New()
	logger.Out = &SliceWriter{}
	h := handlers.MakeBucketHandler(nil, logger)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	h.Handle(w, r)
	bbytes, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte(`{"meta":{"success":true,"message":"not implemented"}}`), bbytes)
}

func TestBucketPostHandlerWorking(t *testing.T) {
	logger := logrus.New()
	logger.Out = &SliceWriter{}
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{},
	}
	h := handlers.MakeBucketHandler(&s, logger)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/buckets/bucket1", strings.NewReader(""))
	h.Handle(w, r)
	bbytes, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte(`{"meta":{"success":true}}`), bbytes)
	assert.NotEmpty(t, s.GetBuckets())
}

func TestBucketDeleteHandlerWorking(t *testing.T) {
	logger := logrus.New()
	logger.Out = &SliceWriter{}
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			"bucket1": map[string][]byte{},
		},
	}
	h := handlers.MakeBucketHandler(&s, logger)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/buckets/bucket1", strings.NewReader(""))
	m := mux.NewRouter()
	m.Handle(h.GetPath(), http.HandlerFunc(h.Handle))
	m.ServeHTTP(w, r)
	bbytes, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte(`{"meta":{"success":true}}`), bbytes)
	assert.Empty(t, s.GetBuckets())
}

type SliceWriter struct{}

func (sr *SliceWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
