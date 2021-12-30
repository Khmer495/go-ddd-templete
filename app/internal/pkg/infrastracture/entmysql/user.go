package entmysql

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util/ccontext"
	"golang.org/x/xerrors"
)

type userRepository struct {
	mysqlClient *ent.Client
}

func NewUserRepository() repository.IUserRepository {
	return userRepository{
		mysqlClient: MysqlClient,
	}
}

func (ur userRepository) Register(ctx context.Context, u entity.User) error {
	err := func() error {
		tx, err := ur.mysqlClient.Tx(ctx)
		if err != nil {
			return xerrors.Errorf("ur.mysqlClient.Tx: %w", err)
		}
		user, err := createUser(ctx, tx.User, u)
		if err != nil {
			return txRollbackAndParseError(tx, err, "createUser")
		}
		_, err = createAuth(ctx, tx.Auth, user.ID)
		if err != nil {
			return txRollbackAndParseError(tx, err, "createAuth")
		}
		return tx.Commit()
	}()
	if err != nil {
		return newTxCommitError(err)
	}
	return nil
}

func (ur userRepository) IsExist(ctx context.Context, userId entity.Id) (bool, error) {
	is, err := isUserExist(ctx, ur.mysqlClient.User, userId)
	if err != nil {
		return false, xerrors.Errorf("isUserExist: %w", err)
	}
	return is, nil
}

func (ur userRepository) Self(ctx context.Context) (entity.User, error) {
	userId, err := ccontext.GetUserId(ctx)
	if err != nil {
		return entity.NilUser, xerrors.Errorf("ccontext.GetUserId: %w", err)
	}
	user, err := findUser(ctx, ur.mysqlClient.User, userId)
	if err != nil {
		return entity.NilUser, xerrors.Errorf("findUserById: %w", err)
	}
	entityUser, err := user.EntityUser()
	if err != nil {
		return entity.NilUser, xerrors.Errorf("user.EntityUser: %w", err)
	}
	return entityUser, nil
}

func (ur userRepository) List(ctx context.Context, limit entity.Limit, page entity.Page) (entity.Users, error) {
	users, err := getUsers(ctx, ur.mysqlClient.User, limit, page, getUsersParam{})
	if err != nil {
		return entity.NilUsers, xerrors.Errorf("getUsers: %w", err)
	}
	entityUsers, err := ent.Users(users).EntityUsers()
	if err != nil {
		return entity.NilUsers, xerrors.Errorf("users.NewEntityUsers: %w", err)
	}
	return entityUsers, nil
}

func (ur userRepository) Select(ctx context.Context, limit entity.Limit, page entity.Page, userIds entity.Ids) (entity.Users, error) {
	users, err := getUsers(ctx, ur.mysqlClient.User, limit, page, getUsersParam{pUserIds: &userIds})
	if err != nil {
		return entity.NilUsers, xerrors.Errorf("getUsers: %w", err)
	}
	entityUsers, err := ent.Users(users).EntityUsers()
	if err != nil {
		return entity.NilUsers, xerrors.Errorf("users.NewEntityUsers: %w", err)
	}
	return entityUsers, nil
}

func (ur userRepository) SearchByNamePrefix(ctx context.Context, limit entity.Limit, page entity.Page, un entity.UserName) (entity.Users, error) {
	users, err := getUsers(ctx, ur.mysqlClient.User, limit, page, getUsersParam{pUserNamePrefix: &un})
	if err != nil {
		return entity.NilUsers, xerrors.Errorf("getUsers: %w", err)
	}
	entityUsers, err := ent.Users(users).EntityUsers()
	if err != nil {
		return entity.NilUsers, xerrors.Errorf("users.NewEntityUsers: %w", err)
	}
	return entityUsers, nil
}

func (ur userRepository) Change(ctx context.Context, u entity.User) error {
	err := func() error {
		tx, err := ur.mysqlClient.Tx(ctx)
		if err != nil {
			return xerrors.Errorf("ur.mysqlClient.Tx: %w", err)
		}
		user, err := findUser(ctx, tx.User, u.Id())
		if err != nil {
			return txRollbackAndParseError(tx, err, "findUser")
		}
		user, err = updateUser(ctx, tx.User, user.ID, u)
		if err != nil {
			return txRollbackAndParseError(tx, err, "updateUser")
		}
		return tx.Commit()
	}()
	if err != nil {
		return newTxCommitError(err)
	}
	return nil
}
