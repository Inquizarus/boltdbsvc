package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/boltdbsvc/handlers"
	"github.com/inquizarus/boltdbsvc/middlewares"
	gorest "github.com/inquizarus/gorest"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := makeLogger(os.Stdout)
	db, err := bolt.Open("svc.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if nil != err {
		logger.Fatal(fmt.Errorf(`Database connection error: %v`, err))
	}
	serveConfig := gorest.ServeConfig{
		Port:        "8080",
		Logger:      logger,
		Middlewares: makeMiddlewares(logger),
		Handlers:    makeHandlers(db, logger),
	}
	start(serveConfig)
}

func start(cfg gorest.ServeConfig) {
	gorest.Serve(cfg)
}

func makeMiddlewares(logger log.StdLogger) []gorest.Middleware {
	return []gorest.Middleware{
		gorest.WithJSONContent(),
		middlewares.WithRequestLogging(logger),
	}
}

func makeLogger(o io.Writer) log.StdLogger {
	logger := log.New()
	logger.SetOutput(o)
	logger.SetFormatter(&log.JSONFormatter{})
	return logger
}

func makeHandlers(db *bolt.DB, logger log.StdLogger) []gorest.Handler {
	return []gorest.Handler{
		handlers.MakeBucketHandler(db, logger),
		handlers.MakeListBucketHandler(db, logger),
	}
}

func errNilOrPanic(err error) {
	if nil != err {
		panic(err)
	}
}
