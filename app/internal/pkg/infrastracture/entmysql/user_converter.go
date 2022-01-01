package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

func entUserTomodelUser(u ent.User) (model.User, error) {
	user, err := model.NewUser(u.Ulid, u.CreatedAt, u.Name)
	if err != nil {
		return model.NilUser, xerrors.Errorf("model.NewUser: %w", err)
	}
	return user, nil
}

func entUsersTomodelUsers(us ent.Users) (model.Users, error) {
	users := model.Users{}
	for _, u := range us {
		if u == nil {
			return model.NilUsers, cerror.NewNotFoundError("u == nil", cerror.InterServerErrorCode.ToString())
		}
		user, err := entUserTomodelUser(*u)
		if err != nil {
			return model.NilUsers, xerrors.Errorf("entUserTomodelUser: %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}
