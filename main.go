package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/boltdbsvc/handlers"
	gorest "github.com/inquizarus/gorest"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {
	configure()
	logger := makeLogger(os.Stdout)
	opts := &bolt.Options{Timeout: 1 * time.Second, ReadOnly: viper.GetBool("reader")}
	db, err := bolt.Open(viper.GetString("db"), 0600, opts)
	if nil != err {
		logger.Fatal(fmt.Errorf(`Database connection error: %v`, err))
	}
	serveConfig := gorest.ServeConfig{
		Port:        viper.GetString("port"),
		Logger:      logger,
		Middlewares: makeMiddlewares(logger),
		Handlers:    makeHandlers(db, logger),
	}
	start(serveConfig)
}

func configure() {
	viper.SetEnvPrefix("boltdbsvc")
	viper.AutomaticEnv()
	viper.SetDefault("port", "8080")
	viper.SetDefault("db", "boltdbsvc.db")
	viper.SetDefault("reader", false)
}

func start(cfg gorest.ServeConfig) {
	gorest.Serve(cfg)
}

func makeMiddlewares(logger log.StdLogger) []gorest.Middleware {
	return []gorest.Middleware{
		gorest.WithJSONContent(),
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
		handlers.MakeItemHandler(db, logger),
	}
}

func errNilOrPanic(err error) {
	if nil != err {
		panic(err)
	}
}
