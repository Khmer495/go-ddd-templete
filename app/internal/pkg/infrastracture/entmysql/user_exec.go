package entmysql

import (
	"context"
	"fmt"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/user"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/Khmer495/go-templete/internal/pkg/util/mysqlerror"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

func createUser(ctx context.Context, uc *ent.UserClient, u entity.User) (ent.User, error) {
	user, err := uc.Create().
		SetUlid(u.Id().Ulid().String()).
		SetCreatedAt(u.CreatedAt().Time()).
		SetUpdatedAt(u.CreatedAt().Time()).
		SetName(u.Name().String()).
		Save(ctx)
	if err != nil {
		mysqlErr := &mysql.MySQLError{}
		if !xerrors.As(err, &mysqlErr) {
			return ent.User{}, cerror.WrapInternalServerError(err, "uc.Create")
		}
		if mysqlErr.Number == mysqlerror.DuplicateErrorNumber {
			return ent.User{}, cerror.Wrap(err, cerror.ErrorLevel, cerror.InvalidArgumentErrorCode, "mysqlErr.Number == mysqlerror.DuplicateErrorNumber", "既に登録されているユーザーです。")
		}
		return ent.User{}, cerror.WrapInternalServerError(err, "uc.Create")
	}
	return *user, nil
}

func isUserExist(ctx context.Context, uc *ent.UserClient, userId entity.Id) (bool, error) {
	uq := uc.Query()
	userQueryFind(uq, userId)
	userQueryActive(uq)
	ok, err := uq.Exist(ctx)
	if err != nil {
		return false, cerror.WrapNotFoundError(err, "uq.Exist", fmt.Sprintf("userId: %sのuserは存在しません。", userId.Ulid().String()))
	}
	return ok, nil
}

func findUserRecord(ctx context.Context, uc *ent.UserClient, userId entity.Id) (ent.User, error) {
	uq := uc.Query()
	userQuerySelect(uq)
	userQueryFind(uq, userId)
	userQueryActive(uq)
	user, err := uq.Only(ctx)
	if err != nil {
		return ent.User{}, parseQueryOnlyByUlid(err, "uq.Only", fmt.Sprintf("userId: %sのuserが存在しません。", userId.Ulid().String()))
	}
	return *user, nil
}

func findUser(ctx context.Context, uc *ent.UserClient, userId entity.Id) (ent.User, error) {
	return findUserRecord(ctx, uc, userId)
}

type getUsersParam struct {
	pUserIds        *entity.Ids
	pUserNamePrefix *entity.UserName
}

func getUsers(ctx context.Context, uc *ent.UserClient, limit entity.Limit, page entity.Page, gup getUsersParam) (ent.Users, error) {
	uq := uc.Query()
	userQuerySelect(uq)
	userQueryActive(uq)
	userQueryOrder(uq)
	userQueryPaging(uq, limit, page)
	if gup.pUserIds != nil {
		uq = uq.Where(user.UlidIn(gup.pUserIds.String()...))
	}
	if gup.pUserNamePrefix != nil {
		uq = uq.Where(user.NameHasPrefix(gup.pUserNamePrefix.String()))
	}
	users, err := uq.All(ctx)
	if err != nil {
		return nil, cerror.WrapInternalServerError(err, "uq.All")
	}
	return users, nil
}

func updateUser(ctx context.Context, uc *ent.UserClient, userPk int, u entity.User) (ent.User, error) {
	user, err := uc.UpdateOneID(userPk).
		SetName(u.Name().String()).
		Save(ctx)
	if err != nil {
		return ent.User{}, cerror.WrapInternalServerError(err, "uc.UpdateOneID.Save")
	}
	return *user, nil
}
