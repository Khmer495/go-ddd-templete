package entmysql

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
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

func (ur userRepository) Register(ctx context.Context, u model.User) error {
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

func (ur userRepository) IsExist(ctx context.Context, userId model.Id) (bool, error) {
	is, err := isUserExist(ctx, ur.mysqlClient.User, userId)
	if err != nil {
		return false, xerrors.Errorf("isUserExist: %w", err)
	}
	return is, nil
}

func (ur userRepository) Self(ctx context.Context) (model.User, error) {
	userId, err := ccontext.GetUserId(ctx)
	if err != nil {
		return model.NilUser, xerrors.Errorf("ccontext.GetUserId: %w", err)
	}
	user, err := findUser(ctx, ur.mysqlClient.User, userId)
	if err != nil {
		return model.NilUser, xerrors.Errorf("findUserById: %w", err)
	}
	modelUser, err := entUserTomodelUser(user)
	if err != nil {
		return model.NilUser, xerrors.Errorf("user.modelUser: %w", err)
	}
	return modelUser, nil
}

func (ur userRepository) List(ctx context.Context, limit model.Limit, page model.Page) (model.Users, error) {
	users, err := getUsers(ctx, ur.mysqlClient.User, limit, page, getUsersParam{})
	if err != nil {
		return model.NilUsers, xerrors.Errorf("getUsers: %w", err)
	}
	modelUsers, err := entUsersTomodelUsers(users)
	if err != nil {
		return model.NilUsers, xerrors.Errorf("users.NewmodelUsers: %w", err)
	}
	return modelUsers, nil
}

func (ur userRepository) Select(ctx context.Context, limit model.Limit, page model.Page, userIds model.Ids) (model.Users, error) {
	users, err := getUsers(ctx, ur.mysqlClient.User, limit, page, getUsersParam{pUserIds: &userIds})
	if err != nil {
		return model.NilUsers, xerrors.Errorf("getUsers: %w", err)
	}
	modelUsers, err := entUsersTomodelUsers(users)
	if err != nil {
		return model.NilUsers, xerrors.Errorf("users.NewmodelUsers: %w", err)
	}
	return modelUsers, nil
}

func (ur userRepository) SearchByNamePrefix(ctx context.Context, limit model.Limit, page model.Page, un model.UserName) (model.Users, error) {
	users, err := getUsers(ctx, ur.mysqlClient.User, limit, page, getUsersParam{pUserNamePrefix: &un})
	if err != nil {
		return model.NilUsers, xerrors.Errorf("getUsers: %w", err)
	}
	modelUsers, err := entUsersTomodelUsers(users)
	if err != nil {
		return model.NilUsers, xerrors.Errorf("users.NewmodelUsers: %w", err)
	}
	return modelUsers, nil
}

func (ur userRepository) Change(ctx context.Context, u model.User) error {
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
