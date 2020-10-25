package app

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/inquizarus/golbag/pkg/handlers"
	"github.com/inquizarus/golbag/pkg/logging"
	"github.com/inquizarus/golbag/pkg/storages"
	"github.com/inquizarus/gorest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func makeServeCmd(log logging.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Starts the built in HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			opts := &bolt.Options{Timeout: 1 * time.Second, ReadOnly: viper.GetBool("reader")}
			db, err := bolt.Open(viper.GetString(databaseVarName), 0600, opts)
			if nil != err {
				log.Fatal(fmt.Errorf(`Database connection error: %v`, err))
			}
			s := storages.MakeBoltDBStorage(db)
			serveConfig := gorest.ServeConfig{
				Port:   viper.GetString(portVarName),
				Logger: log,
				Middlewares: []gorest.Middleware{
					gorest.WithJSONContent(),
				},
				Handlers: []gorest.Handler{
					handlers.MakeBucketHandler(s, log),
					handlers.MakeListBucketHandler(s, log),
					handlers.MakeItemHandler(s, log),
				},
			}
			gorest.Serve(serveConfig)
		},
	}
	command.Flags().StringVarP(&port, portVarName, portVarShortName, portVarDefaultValue, "Define which port should be used.")
	viper.BindPFlag(portVarName, command.Flags().Lookup(portVarName))
	return command
}
