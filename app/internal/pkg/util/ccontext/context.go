package ccontext

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

type firebaseUserIdKey struct{}

func SetFirebaseUserId(ctx context.Context, val string) context.Context {
	ctx = context.WithValue(ctx, firebaseUserIdKey{}, val)
	return ctx
}

func GetFirebaseUserId(ctx context.Context) (string, error) {
	val := ctx.Value(firebaseUserIdKey{})
	if val == nil {
		return "", cerror.New(cerror.ErrorLevel, cerror.NotFoundErrorCode, "ctx.Value", cerror.InterServerErrorCode.ToString())
	}
	valString, ok := val.(string)
	if !ok {
		return "", cerror.New(cerror.ErrorLevel, cerror.InterServerErrorCode, "ctx.Value.(string)", cerror.InterServerErrorCode.ToString())
	}
	return valString, nil
}

type userIdKey struct{}

func SetUserId(ctx context.Context, val model.Id) context.Context {
	ctx = context.WithValue(ctx, userIdKey{}, val.Ulid().String())
	return ctx
}

func GetUserId(ctx context.Context) (model.Id, error) {
	val := ctx.Value(userIdKey{})
	if val == nil {
		return model.NilId, cerror.New(cerror.ErrorLevel, cerror.NotFoundErrorCode, "ctx.Value", cerror.InterServerErrorCode.ToString())
	}
	valString, ok := val.(string)
	if !ok {
		return model.NilId, cerror.New(cerror.ErrorLevel, cerror.InterServerErrorCode, "ctx.Value.(string)", cerror.InterServerErrorCode.ToString())
	}
	userId, err := model.NewId(valString)
	if err != nil {
		return model.NilId, xerrors.Errorf("model.NewId: %w", err)
	}
	return userId, nil
}
