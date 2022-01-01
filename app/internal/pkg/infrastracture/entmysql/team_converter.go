package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

func entTeamTomodelTeam(t ent.Team) (model.Team, error) {
	pCreateUser := t.Edges.CreateUser
	if pCreateUser == nil {
		return model.NilTeam, cerror.NewInternalServerError("pCreateUser == nil")
	}
	createUser, err := entUserTomodelUser(*pCreateUser)
	if err != nil {
		return model.NilTeam, xerrors.Errorf("t.Edges.CreateUser.modelUser: %w", err)
	}
	users, err := entTeamUsersTomodelUsers(t.Edges.TeamUsers)
	if err != nil {
		return model.NilTeam, xerrors.Errorf("entTeamUsersTomodelUsers: %w", err)
	}
	team, err := model.NewTeam(t.Ulid, t.CreatedAt, createUser, t.Name, t.Description, users)
	if err != nil {
		return model.NilTeam, xerrors.Errorf("model.NewTeam: %w", err)
	}
	return team, err
}

func entTeamsTomodelTeams(ts ent.Teams) (model.Teams, error) {
	teams := model.Teams{}
	for _, t := range ts {
		if t == nil {
			return model.NilTeams, cerror.NewNotFoundError("t == nil", cerror.InterServerErrorCode.ToString())
		}
		team, err := entTeamTomodelTeam(*t)
		if err != nil {
			return model.NilTeams, xerrors.Errorf("t.modelTeam: %w", err)
		}
		teams = append(teams, &team)
	}
	return teams, nil
}
