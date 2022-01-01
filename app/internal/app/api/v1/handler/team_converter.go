package handler

import (
	"context"

	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"golang.org/x/xerrors"
)

func frommodelTeamToOpenapiTeam(ctx context.Context, t model.Team) (openapi.Team, error) {
	users, err := frommodelUsersToOpenapiUsers(ctx, t.Users())
	if err != nil {
		return openapi.Team{}, xerrors.Errorf("frommodelUsersToOpenapiUsers: %w", err)
	}
	return openapi.Team{
		Id:          t.Id().Ulid().String(),
		Name:        t.Name().String(),
		Description: t.Description().String(),
		Users:       users,
	}, nil
}

func frommodelTeamsToOpenapiTeams(ctx context.Context, ts model.Teams) ([]openapi.Team, error) {
	teams := []openapi.Team{}
	for _, t := range ts {
		team, err := frommodelTeamToOpenapiTeam(ctx, *t)
		if err != nil {
			return nil, xerrors.Errorf("frommodelTeamToOpenapiTeam: %w", err)
		}
		teams = append(teams, team)
	}
	return teams, nil
}
