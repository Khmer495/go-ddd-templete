package entmysql

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/auth"
	"github.com/Khmer495/go-templete/internal/pkg/util/ccontext"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

type authRepository struct {
	mysqlClient *ent.Client
}

func NewAuthRepository() repository.IAuthRepository {
	return authRepository{
		mysqlClient: MysqlClient,
	}
}

func (ar authRepository) FindUserPkByFirebaseUserId(ctx context.Context) (int, error) {
	firebaseUserId, err := ccontext.GetFirebaseUserId(ctx)
	if err != nil {
		return 0, xerrors.Errorf("ccontext.GetFirebaseUserId: %w", err)
	}
	auth, err := ar.mysqlClient.Auth.Query().
		Where(auth.FirebaseUserIDEQ(firebaseUserId)).
		First(ctx)
	if err != nil {
		entNotFountError := &ent.NotFoundError{}
		if xerrors.As(err, &entNotFountError) {
			return 0, cerror.Wrap(err, cerror.ErrorLevel, cerror.NotFoundErrorCode, "ar.mysqlClient.Auth.Query", "存在しないユーザーです。")
		}
		return 0, cerror.Wrap(err, cerror.ErrorLevel, cerror.InterServerErrorCode, "ar.mysqlClient.Auth.Query", cerror.InterServerErrorCode.ToString())
	}
	return auth.UserID, nil
}

func (ar authRepository) FindUserIdByFirebaseUserId(ctx context.Context) (model.Id, error) {
	firebaseUserId, err := ccontext.GetFirebaseUserId(ctx)
	if err != nil {
		return model.NilId, xerrors.Errorf("ccontext.GetFirebaseUserId: %w", err)
	}
	auth, err := ar.mysqlClient.Auth.Query().
		Where(auth.FirebaseUserIDEQ(firebaseUserId)).
		WithUser(func(uq *ent.UserQuery) {
			uq.Select("ulid")
		}).
		First(ctx)
	if err != nil {
		entNotFountError := &ent.NotFoundError{}
		if xerrors.As(err, &entNotFountError) {
			return model.NilId, cerror.Wrap(err, cerror.ErrorLevel, cerror.NotFoundErrorCode, "ar.mysqlClient.Auth.Query", "存在しないユーザーです。")
		}
		return model.NilId, cerror.Wrap(err, cerror.ErrorLevel, cerror.InterServerErrorCode, "ar.mysqlClient.Auth.Query", cerror.InterServerErrorCode.ToString())
	}
	userId, err := model.NewId(auth.Edges.User.Ulid)
	if err != nil {
		return model.NilId, xerrors.Errorf("model.NewId: %w", err)
	}
	return userId, nil
}

func createAuth(ctx context.Context, ac *ent.AuthClient, userPk int) (ent.Auth, error) {
	firebaseAuthId, err := ccontext.GetFirebaseUserId(ctx)
	if err != nil {
		return ent.Auth{}, xerrors.Errorf("ccontext.GetFirebaseUserId: %w", err)
	}
	auth, err := ac.Create().
		SetUserID(userPk).
		SetFirebaseUserID(firebaseAuthId).
		Save(ctx)
	if err != nil {
		return ent.Auth{}, cerror.Wrap(err, cerror.ErrorLevel, cerror.InterServerErrorCode, "ac.Create", cerror.InterServerErrorCode.ToString())
	}
	return *auth, nil
}
