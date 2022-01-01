//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"

	model "github.com/Khmer495/go-templete/internal/pkg/domain/model"
)

type ITeamRepository interface {
	Create(ctx context.Context, t model.Team) error
	IsExist(ctx context.Context, teamId model.Id) (bool, error)
	One(ctx context.Context, teamId model.Id) (model.Team, error)
	List(ctx context.Context, limit model.Limit, page model.Page) (model.Teams, error)
	SearchByNamePrefix(ctx context.Context, limit model.Limit, page model.Page, un model.TeamName) (model.Teams, error)
	Change(ctx context.Context, t model.Team) error
	Delete(ctx context.Context, teamId model.Id) error
	Join(ctx context.Context, teamId model.Id, userId model.Id) error
}
