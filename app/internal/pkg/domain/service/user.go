//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package service

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/util/ccontext"
	"golang.org/x/xerrors"
)

type IUserService interface {
	SelfId(ctx context.Context) (model.Id, error)
	IsSelf(ctx context.Context, userId model.Id) (bool, error)
}

type userService struct{}

func NewUserService() IUserService {
	return userService{}
}

func (us userService) SelfId(ctx context.Context) (model.Id, error) {
	selfId, err := ccontext.GetUserId(ctx)
	if err != nil {
		return model.NilId, xerrors.Errorf("ccontext.GetUserId: %w", err)
	}
	return selfId, nil
}

func (us userService) IsSelf(ctx context.Context, userId model.Id) (bool, error) {
	selfUserId, err := ccontext.GetUserId(ctx)
	if err != nil {
		return false, xerrors.Errorf("ccontext.GetUserId: %w", err)
	}
	if selfUserId != userId {
		return false, nil
	}
	return true, nil
}
