//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package firebase

import (
	"context"

	fb "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type IFirebaseClient interface {
	VerifyIdToken(ctx context.Context, idToken string) (*auth.Token, error)
	GetUid(token *auth.Token) string
}

type firebaseClient struct {
	app             *fb.App
	authClient      *auth.Client
	messagingClient *messaging.Client
}

func NewFirebaseClient() IFirebaseClient {
	ctx := context.Background()

	opt := option.WithCredentialsFile("some_credentials.json")
	app, err := fb.NewApp(ctx, nil, opt)
	if err != nil {
		zap.S().Panicf("firebase.NewApp: %+v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		zap.S().Panicf("App.Auth: %+v", err)
	}

	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		zap.S().Panicf("App.Messaging: %w", err)
	}

	return firebaseClient{
		app:             app,
		authClient:      authClient,
		messagingClient: messagingClient,
	}
}
