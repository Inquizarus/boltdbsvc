package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/inquizarus/golbag/pkg/handlers"
	"github.com/inquizarus/golbag/pkg/storages"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestBucketsHandlerWorking(t *testing.T) {
	logger := logrus.New()
	logger.Out = &SliceWriter{}
	s := storages.MapStorage{
		Map: map[string]map[string][]byte{
			"bucket1": {},
		},
	}
	h := handlers.MakeListBucketHandler(&s, logger)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/buckets", strings.NewReader(""))
	h.Handle(w, r)
	bbytes, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, []byte(`{"meta":{"success":true},"data":["bucket1"]}`), bbytes)
}
