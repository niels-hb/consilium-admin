package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"github.com/niels-hb/consilium-admin/handlers"
	"github.com/thatisuday/commando"
	"google.golang.org/api/option"
)

func main() {
	initializeApp()
	initializeCommando()
}

func initializeApp() (*firebase.App, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("service_account.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore: %v", err)
	}

	handlers.Context = ctx
	handlers.Client = client

	return app, err
}

func initializeCommando() {
	commando.
		SetExecutableName("consilium-admin").
		SetVersion("1.0.0").
		SetDescription("This tool contains administrative commands for the Consilium backend.")

	commando.
		Register("export").
		SetShortDescription("export a users documents to a JSON file").
		AddFlag("uid,u", "UID of the user", commando.String, nil).
		AddFlag("target,t", "target file", commando.String, nil).
		SetAction(handlers.Export)

	commando.
		Register("import").
		SetShortDescription("import data from a JSON file").
		AddFlag("uid,u", "UID of the user", commando.String, nil).
		AddFlag("source,s", "source file", commando.String, nil).
		AddFlag("dry-run", "don't send actual requests to the server", commando.Bool, nil).
		SetAction(handlers.Import)

	commando.
		Register("migrate").
		SetShortDescription("migrate documents to a different user").
		AddFlag("from,f", "UID of the current owner", commando.String, nil).
		AddFlag("to,t", "UID of the user to which ownership should be transferred", commando.String, nil).
		AddFlag("dry-run", "don't send actual requests to the server", commando.Bool, nil).
		SetAction(handlers.Migrate)

	commando.Parse(nil)
}
