package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

func entTeamUserTomodelUser(tu ent.TeamUser) (model.User, error) {
	pUser := tu.Edges.User
	if pUser == nil {
		return model.NilUser, cerror.NewNotFoundError("pUser == nil", cerror.InterServerErrorCode.ToString())
	}
	user, err := entUserTomodelUser(*pUser)
	if err != nil {
		return model.NilUser, xerrors.Errorf("entUserTomodelUser: %w", err)
	}
	return user, nil
}

func entTeamUsersTomodelUsers(tus ent.TeamUsers) (model.Users, error) {
	users := model.Users{}
	for _, tu := range tus {
		if tu == nil {
			return model.NilUsers, cerror.NewNotFoundError("tu == nil", cerror.InterServerErrorCode.ToString())
		}
		user, err := entTeamUserTomodelUser(*tu)
		if err != nil {
			return model.NilUsers, xerrors.Errorf("entTeamUserTomodelUser: %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}
