package main

import (
	"context"
	"go-http/conf"
	"go-http/internal/api/v1/modules/cron_module"
	"go-http/pkg/app"
	"log"

	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	err := supertokens.Init(conf.SuperTokensConfig)
	if err != nil {
		panic(err.Error())
	}

	app, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer app.Exit(context.Background())

	cron_module.NewCronInstance()
	NewServer(app).Start()
}
