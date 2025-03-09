package config

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitFirebase(env Env) (*firestore.Client, *auth.Client, error) {
	ctx := context.Background()
	jsonCred := env.GoogleApplicationCredentials

	conf := &firebase.Config{
		ProjectID: env.FirestoreProjectID,
	}

	app, err := firebase.NewApp(ctx, conf, option.WithCredentialsJSON([]byte(jsonCred)))
	if err != nil {
		return nil, nil, err
	}

	fsClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, nil, err
	}

	return fsClient, authClient, nil
}
