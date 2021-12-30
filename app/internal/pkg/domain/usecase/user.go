//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

type IUserUsecase interface {
	Register(ctx context.Context, name string) (entity.User, error)

	GetSelfProfile(ctx context.Context) (entity.User, error)

	GetList(ctx context.Context, limit int, page int, pName *string) (entity.Users, error)

	ChangeSelfProfile(ctx context.Context, pName *string) (entity.User, error)
}

type userUsecase struct {
	us service.IUserService
	ur repository.IUserRepository
}

func NewUserUsecase(
	us service.IUserService,
	ur repository.IUserRepository,
) IUserUsecase {
	return userUsecase{
		us: us,
		ur: ur,
	}
}

func (uu userUsecase) Register(ctx context.Context, name string) (entity.User, error) {
	_, err := uu.us.SelfId(ctx)
	// err == nil、すなわちuserIdが取得できた場合、既に登録済みのユーザーであるため、新規登録を拒否する
	if err == nil {
		return entity.NilUser, cerror.Wrap(err, cerror.ErrorLevel, cerror.AlreadyRegisterdUserErrorCode, "ccontext.GetUserId", "既に登録されているユーザーです。")
	}
	// ユーザーが登録されていないことを保証するエラーはNotFoundのみである
	if !cerror.IsCode(err, cerror.NotFoundErrorCode) {
		return entity.NilUser, cerror.WrapInternalServerError(err, "ccontext.GetUserPk")
	}

	user, err := entity.InitUser(name)
	if err != nil {
		return entity.NilUser, parseEntityConstractorError(err, "entity.InitUser")
	}

	err = uu.ur.Register(ctx, user)
	if err != nil {
		return entity.NilUser, cerror.WrapInternalServerError(err, "uu.ur.Create")
	}

	return user, nil
}

func (uu userUsecase) GetSelfProfile(ctx context.Context) (entity.User, error) {
	user, err := uu.ur.Self(ctx)
	if err != nil {
		return entity.NilUser, cerror.WrapInternalServerError(err, "uu.ur.GetSelf")
	}
	return user, nil
}

func (uu userUsecase) GetList(ctx context.Context, limitInt int, pageInt int, pNameString *string) (entity.Users, error) {
	limit, err := entity.NewLimit(limitInt)
	if err != nil {
		return entity.NilUsers, parseEntityConstractorError(err, "entity.NewLimit")
	}
	page, err := entity.NewPage(pageInt)
	if err != nil {
		return entity.NilUsers, parseEntityConstractorError(err, "entity.NewPage")
	}

	var users entity.Users
	if pNameString == nil {
		users, err = uu.ur.List(ctx, limit, page)
		if err != nil {
			return entity.NilUsers, cerror.WrapInternalServerError(err, "uu.ur.GetList")
		}
	} else {
		name, err := entity.NewUserName(*pNameString)
		if err != nil {
			return entity.NilUsers, parseEntityConstractorError(err, "entity.NewUserName")
		}
		users, err = uu.ur.SearchByNamePrefix(ctx, limit, page, name)
		if err != nil {
			return entity.NilUsers, cerror.WrapInternalServerError(err, "uu.ur.SearchByNamePrefix")
		}
	}
	return users, nil
}

func (uu userUsecase) ChangeSelfProfile(ctx context.Context, pName *string) (entity.User, error) {
	user, err := uu.ur.Self(ctx)
	if err != nil {
		return entity.NilUser, cerror.WrapInternalServerError(err, "uu.ur.GetSelf")
	}
	if pName != nil {
		if err := user.SetName(*pName); err != nil {
			return entity.NilUser, parseEntityConstractorError(err, "user.SetName: %w")
		}
	}
	if err := uu.ur.Change(ctx, user); err != nil {
		return entity.NilUser, xerrors.Errorf("uu.ur.Change: %w", err)
	}
	return user, nil
}
