package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/golbag/handlers"
	"github.com/inquizarus/golbag/storages"
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
	s := storages.MakeBoltDBStorage(db)
	serveConfig := gorest.ServeConfig{
		Port:        viper.GetString("port"),
		Logger:      logger,
		Middlewares: makeMiddlewares(logger),
		Handlers:    makeHandlers(s, logger),
	}
	start(serveConfig)
}

func configure() {
	viper.SetEnvPrefix("golbag")
	viper.AutomaticEnv()
	viper.SetDefault("port", "8080")
	viper.SetDefault("db", "golbag.db")
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

func makeHandlers(s storages.Storage, logger log.StdLogger) []gorest.Handler {
	return []gorest.Handler{
		handlers.MakeBucketHandler(s, logger),
		handlers.MakeListBucketHandler(s, logger),
		handlers.MakeItemHandler(s, logger),
	}
}

func errNilOrPanic(err error) {
	if nil != err {
		panic(err)
	}
}
