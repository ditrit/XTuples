// Package app groups all of the application's environment configuration and
// database connection dependencies and provides a single entry point for
// initializing the application.
package app

import (
	"context"
	"database/sql"
	"fmt"
	"go-http/pkg/settings"
	"go-http/pkg/settings/cli"
	"go-http/pkg/settings/database"
)

// App is the application struct which contains the environment configuration
// and the database connection.
type App struct {
	conf *settings.EnvConfig // conf is the environment configuration
	db   *sql.DB             // db is the database connection
}

// Conf returns the environment configuration.
func (a *App) Conf() *settings.EnvConfig {
	return a.conf
}

// Db returns the database connection.
func (a *App) Db() *sql.DB {
	return a.db
}

// NewApp setups the environment for the go app on the initialization
// and returns the App struct which contains the read environment
// configuration and the database connection.
//
// Example:
//
//	app, err := app.NewApp(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer app.Exit(context.Background())
func NewApp(ctx context.Context, cliArgs ...cli.Args) (*App, error) {
	args := cli.NewArgs(cliArgs...)

	fmt.Printf("Initializing app with the %v config file\n", args.Conf)

	conf, err := settings.NewEnvConfig(args.Conf)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file, %v", err)
	}

	dbConf := conf.Database
	dsnString := database.DsnString(dbConf.Host, dbConf.User, dbConf.Pass, dbConf.Name, dbConf.Port, dbConf.SSL, dbConf.Timezone)
	db, err := database.NewDbConn(dsnString)
	if err != nil {
		return nil, fmt.Errorf("error creating database connection, %v", err)
	}

	return &App{conf, db}, nil
}

// Exit function groups all the exit functions of the app (e.g. closing the database connection)
// and returns an error if any of the exit functions fail.
func (a *App) Exit(ctx context.Context) error {
	// Close the database connection
	if err := a.Db().Close(); err != nil {
		return fmt.Errorf("failed to close db conn: %v", err)
	}
	return nil
}
