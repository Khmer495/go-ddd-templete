package ent

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

func (u *User) EntityUser() (entity.User, error) {
	if u == nil {
		return entity.NilUser, cerror.NewNotFoundError("u == nil", cerror.InterServerErrorCode.ToString())
	}
	user, err := entity.NewUser(u.Ulid, u.CreatedAt, u.Name)
	if err != nil {
		return entity.NilUser, xerrors.Errorf("entity.NewUser: %w", err)
	}
	return user, nil
}

func (us Users) EntityUsers() (entity.Users, error) {
	users := entity.Users{}
	for _, u := range us {
		user, err := u.EntityUser()
		if err != nil {
			if cerror.IsCode(err, cerror.NotFoundErrorCode) {
				continue
			}
			return entity.NilUsers, xerrors.Errorf("u.EntityUser: %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}
