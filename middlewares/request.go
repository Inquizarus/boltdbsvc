package middlewares

import (
	"net/http"

	"github.com/inquizarus/gorest"
	"github.com/sirupsen/logrus"
)

// WithRequestLogging handles logging of all incoming requests
func WithRequestLogging(logger logrus.StdLogger) gorest.Middleware {
	return func(f http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logger.Printf(`incoming %s request`, r.Method)
			f.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
