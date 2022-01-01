//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

type IUserUsecase interface {
	Register(ctx context.Context, name string) (model.User, error)

	GetSelfProfile(ctx context.Context) (model.User, error)

	GetList(ctx context.Context, limit int, page int, pName *string) (model.Users, error)

	ChangeSelfProfile(ctx context.Context, pName *string) (model.User, error)
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

func (uu userUsecase) Register(ctx context.Context, name string) (model.User, error) {
	_, err := uu.us.SelfId(ctx)
	// err == nil、すなわちuserIdが取得できた場合、既に登録済みのユーザーであるため、新規登録を拒否する
	if err == nil {
		return model.NilUser, cerror.Wrap(err, cerror.ErrorLevel, cerror.AlreadyRegisterdUserErrorCode, "ccontext.GetUserId", "既に登録されているユーザーです。")
	}
	// ユーザーが登録されていないことを保証するエラーはNotFoundのみである
	if !cerror.IsCode(err, cerror.NotFoundErrorCode) {
		return model.NilUser, cerror.WrapInternalServerError(err, "ccontext.GetUserPk")
	}

	user, err := model.InitUser(name)
	if err != nil {
		return model.NilUser, parsemodelConstractorError(err, "model.InitUser")
	}

	err = uu.ur.Register(ctx, user)
	if err != nil {
		return model.NilUser, cerror.WrapInternalServerError(err, "uu.ur.Create")
	}

	return user, nil
}

func (uu userUsecase) GetSelfProfile(ctx context.Context) (model.User, error) {
	user, err := uu.ur.Self(ctx)
	if err != nil {
		return model.NilUser, cerror.WrapInternalServerError(err, "uu.ur.GetSelf")
	}
	return user, nil
}

func (uu userUsecase) GetList(ctx context.Context, limitInt int, pageInt int, pNameString *string) (model.Users, error) {
	limit, err := model.NewLimit(limitInt)
	if err != nil {
		return model.NilUsers, parsemodelConstractorError(err, "model.NewLimit")
	}
	page, err := model.NewPage(pageInt)
	if err != nil {
		return model.NilUsers, parsemodelConstractorError(err, "model.NewPage")
	}

	var users model.Users
	if pNameString == nil {
		users, err = uu.ur.List(ctx, limit, page)
		if err != nil {
			return model.NilUsers, cerror.WrapInternalServerError(err, "uu.ur.GetList")
		}
	} else {
		name, err := model.NewUserName(*pNameString)
		if err != nil {
			return model.NilUsers, parsemodelConstractorError(err, "model.NewUserName")
		}
		users, err = uu.ur.SearchByNamePrefix(ctx, limit, page, name)
		if err != nil {
			return model.NilUsers, cerror.WrapInternalServerError(err, "uu.ur.SearchByNamePrefix")
		}
	}
	return users, nil
}

func (uu userUsecase) ChangeSelfProfile(ctx context.Context, pName *string) (model.User, error) {
	user, err := uu.ur.Self(ctx)
	if err != nil {
		return model.NilUser, cerror.WrapInternalServerError(err, "uu.ur.GetSelf")
	}
	if pName != nil {
		if err := user.SetName(*pName); err != nil {
			return model.NilUser, parsemodelConstractorError(err, "user.SetName: %w")
		}
	}
	if err := uu.ur.Change(ctx, user); err != nil {
		return model.NilUser, xerrors.Errorf("uu.ur.Change: %w", err)
	}
	return user, nil
}
