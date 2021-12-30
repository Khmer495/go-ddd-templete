package ent

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

func (t *Team) EntityTeam() (entity.Team, error) {
	if t == nil {
		return entity.NilTeam, cerror.NewNotFoundError("t == nil", cerror.InterServerErrorCode.ToString())
	}
	createUser, err := t.Edges.CreateUser.EntityUser()
	if err != nil {
		return entity.NilTeam, xerrors.Errorf("t.Edges.CreateUser.EntityUser: %w", err)
	}
	users, err := TeamUsers(t.Edges.TeamUsers).EntityUsers()
	if err != nil {
		return entity.NilTeam, xerrors.Errorf("t.Edges.TeamUsers.EntityUsers: %w", err)
	}
	team, err := entity.NewTeam(t.Ulid, t.CreatedAt, createUser, t.Name, t.Description, users)
	if err != nil {
		return entity.NilTeam, xerrors.Errorf("entity.NewTeam: %w", err)
	}
	return team, err
}

func (ts Teams) EntityTeams() (entity.Teams, error) {
	teams := entity.Teams{}
	for _, t := range ts {
		team, err := t.EntityTeam()
		if err != nil {
			if cerror.IsCode(err, cerror.NotFoundErrorCode) {
				continue
			}
			return entity.NilTeams, xerrors.Errorf("t.EntityTeam: %w", err)
		}
		teams = append(teams, &team)
	}
	return teams, nil
}
