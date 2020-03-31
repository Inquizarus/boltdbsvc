package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/inquizarus/boltdbsvc/handlers"
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

type SliceWriter struct{}

func (sr *SliceWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
