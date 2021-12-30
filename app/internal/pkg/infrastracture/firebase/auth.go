package firebase

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"go.uber.org/zap"
)

func (fc firebaseClient) VerifyIdToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := fc.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		zap.S().Errorf("firebaseClient.VerifyIDToken: %+v", err)
		return nil, cerror.Wrap(err, cerror.ErrorLevel, cerror.AuthenticationErrorCode, "client.VerifyIDToken", "Invalid IdToken.")
	}
	return token, nil
}

func (fc firebaseClient) GetUid(token *auth.Token) string {
	return token.UID
}
