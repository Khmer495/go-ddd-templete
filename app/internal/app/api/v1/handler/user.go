package handler

import (
	"context"
	"net/http"

	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/usecase"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type userHandler struct {
	uu usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) IUserHandler {
	return userHandler{
		uu: uu,
	}
}

func fromEntityUserToOpenapiUser(ctx context.Context, u entity.User) (openapi.User, error) {
	return openapi.User{
		Id:   u.Id().Ulid().String(),
		Name: u.Name().String(),
	}, nil
}

func fromEntityUsersToOpenapiUsers(ctx context.Context, us entity.Users) ([]openapi.User, error) {
	users := []openapi.User{}
	for _, u := range us {
		user, err := fromEntityUserToOpenapiUser(ctx, *u)
		if err != nil {
			return nil, xerrors.Errorf("fromEntityUserToOpenapiUser: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (uh userHandler) GetLogin(ctx echo.Context) error {
	user, err := uh.uu.GetSelfProfile(ctx.Request().Context())
	if err != nil {
		return NewInternalServerError(ctx)
	}
	res, err := fromEntityUserToOpenapiUser(ctx.Request().Context(), user)
	if err != nil {
		zap.S().Errorf("fromEntityUserToOpenapiUser: %+v", err)
		return NewInternalServerError(ctx)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (uh userHandler) GetProfile(ctx echo.Context) error {
	user, err := uh.uu.GetSelfProfile(ctx.Request().Context())
	if err != nil {
		zap.S().Errorf("uh.uu.GetSelfProfile: %+v", err)
		return NewInternalServerError(ctx)
	}
	res, err := fromEntityUserToOpenapiUser(ctx.Request().Context(), user)
	if err != nil {
		zap.S().Errorf("fromEntityUserToOpenapiUser: %+v", err)
		return NewInternalServerError(ctx)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (uh userHandler) PutProfile(ctx echo.Context) error {
	req := openapi.PutProfileJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		zap.S().Errorf("ctx.Bind: %+v", err)
		return NewInvalidRequestFormatError(ctx, err)
	}
	user, err := uh.uu.ChangeSelfProfile(ctx.Request().Context(), req.Name)
	if err != nil {
		zap.S().Errorf("uh.uu.ChangeSelfProfile: %+v", err)
		if cerror.IsCode(err, cerror.InvalidArgumentErrorCode) {
			return NewInvalidArgumentError(ctx, cerror.As(err).ClientMsg())
		}
		return NewInternalServerError(ctx)
	}
	res, err := fromEntityUserToOpenapiUser(ctx.Request().Context(), user)
	if err != nil {
		zap.S().Errorf("fromEntityUserToOpenapiUser: %+v", err)
		return NewInternalServerError(ctx)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (uh userHandler) PostUser(ctx echo.Context) error {
	req := openapi.PostUserJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		zap.S().Errorf("ctx.Bind: %+v", err)
		return NewInvalidRequestFormatError(ctx, err)
	}
	user, err := uh.uu.Register(ctx.Request().Context(), req.Name)
	if err != nil {
		zap.S().Errorf("uh.uu.Register: %+v", err)
		if cerror.IsCode(err, cerror.AlreadyRegisterdUserErrorCode) {
			return NewAlreadyRegisteredUserError(ctx)
		}
		if cerror.IsCode(err, cerror.InvalidArgumentErrorCode) {
			return NewInvalidArgumentError(ctx, cerror.As(err).ClientMsg())
		}
		return NewInternalServerError(ctx)
	}
	res, err := fromEntityUserToOpenapiUser(ctx.Request().Context(), user)
	if err != nil {
		zap.S().Errorf("fromEntityUserToOpenapiUser: %+v", err)
		return NewInternalServerError(ctx)
	}
	return ctx.JSON(http.StatusCreated, res)
}
