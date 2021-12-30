package echo

import (
	"github.com/Khmer495/go-templete/internal/app/api/v1/handler"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/firebase"
	"github.com/Khmer495/go-templete/internal/pkg/util/ccontext"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func NewFirebaseAuth(ar repository.IAuthRepository, fc firebase.IFirebaseClient) middleware.KeyAuthValidator {
	return func(idToken string, ctx echo.Context) (bool, error) {
		token, err := fc.VerifyIdToken(ctx.Request().Context(), idToken)
		if err != nil {
			zap.S().Errorf("client.VerifyIDToken: %+v", err)
			// return false, cerror.Wrap(err, cerror.ErrorLevel, cerror.InvalidTokenErrorCode, "client.VerifyIDToken", "不正なトークンです。")
			return false, handler.NewInvalidTokenError(ctx)
		}
		firebaseUserId := fc.GetUid(token)
		ctx.SetRequest(ctx.Request().WithContext(ccontext.SetFirebaseUserId(ctx.Request().Context(), firebaseUserId)))
		zap.S().Debugf("firebaseUserId : %s\n", firebaseUserId)

		userId, err := ar.FindUserIdByFirebaseUserId(ctx.Request().Context())
		if err != nil {
			if !cerror.IsCode(err, cerror.NotFoundErrorCode) {
				zap.S().Errorf("ar.FindUserIdByFirebaseUserId: %+v", err)
				// return false, xerrors.Errorf("if !cerror.IsCode(cerror.NotFoundErrorCode): %w", err)
				return false, handler.NewInternalServerError(ctx)
			}
			// Error Code is cerror.NotFoundErrorCode
			if ctx.Request().Method == "POST" && ctx.Path() == "/api/v1/users" { // TODO: バージョンに依存しない書き方
				return true, nil
			}
			zap.S().Errorf("ar.FindUserIdByFirebaseUserId: %+v", err)
			// return false, cerror.Wrap(err, cerror.ErrorLevel, cerror.UnregisterdUserErrorCode, `if ctx.Request().Method == "POST" && ctx.Path() == "/users"`, "未登録のユーザーです。")
			return false, handler.NewNotRegisteredUserError(ctx)
		}
		ctx.SetRequest(ctx.Request().WithContext(ccontext.SetUserId(ctx.Request().Context(), userId)))
		zap.S().Debugf("userId: %d", userId)

		return true, nil
	}
}

func NewSimpleAuth(ar repository.IAuthRepository) middleware.KeyAuthValidator {
	return func(idToken string, ctx echo.Context) (bool, error) {
		firebaseUserId := idToken
		ctx.SetRequest(ctx.Request().WithContext(ccontext.SetFirebaseUserId(ctx.Request().Context(), firebaseUserId)))
		zap.S().Debugf("firebaseUserId: %s", firebaseUserId)

		userId, err := ar.FindUserIdByFirebaseUserId(ctx.Request().Context())
		if err != nil {
			if !cerror.IsCode(err, cerror.NotFoundErrorCode) {
				zap.S().Errorf("ar.FindUserIdByFirebaseUserId: %+v", err)
				// return false, xerrors.Errorf("if !cerror.IsCode(cerror.NotFoundErrorCode): %w", err)
				return false, handler.NewInternalServerError(ctx)
			}
			// Error Code is cerror.NotFoundErrorCode
			if ctx.Request().Method == "POST" && ctx.Path() == "/api/v1/users" { // TODO: バージョンに依存しない書き方
				return true, nil
			}
			zap.S().Errorf("ar.FindUserIdByFirebaseUserId: %+v", err)
			// return false, cerror.Wrap(err, cerror.ErrorLevel, cerror.UnauthorizedErrorCode, `if ctx.Request().Method == "POST" && ctx.Path() == "/users"`, "未登録のユーザーです。")
			return false, handler.NewNotRegisteredUserError(ctx)
		}
		ctx.SetRequest(ctx.Request().WithContext(ccontext.SetUserId(ctx.Request().Context(), userId)))
		zap.S().Debugf("userId: %d", userId)

		return true, nil
	}
}

func NewNoAuth() middleware.KeyAuthValidator {
	return func(string, echo.Context) (bool, error) {
		return true, nil
	}
}
