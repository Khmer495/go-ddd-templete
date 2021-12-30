package ent

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

func (tu *TeamUser) EntityUser() (entity.User, error) {
	if tu == nil {
		return entity.NilUser, cerror.NewNotFoundError("tu == nil", cerror.InterServerErrorCode.ToString())
	}
	user, err := tu.Edges.User.EntityUser()
	if err != nil {
		return entity.NilUser, xerrors.Errorf("tu.Edges.User.EntityUser: %w", err)
	}
	return user, nil
}

func (tus TeamUsers) EntityUsers() (entity.Users, error) {
	users := entity.Users{}
	for _, tu := range tus {
		user, err := tu.EntityUser()
		if err != nil {
			if cerror.IsCode(err, cerror.NotFoundErrorCode) {
				continue
			}
			return entity.NilUsers, xerrors.Errorf("tu.EntityUser: %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}
