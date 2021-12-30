//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"

	entity "github.com/Khmer495/go-templete/internal/pkg/domain/entity"
)

type ITeamRepository interface {
	Create(ctx context.Context, t entity.Team) error
	IsExist(ctx context.Context, teamId entity.Id) (bool, error)
	One(ctx context.Context, teamId entity.Id) (entity.Team, error)
	List(ctx context.Context, limit entity.Limit, page entity.Page) (entity.Teams, error)
	SearchByNamePrefix(ctx context.Context, limit entity.Limit, page entity.Page, un entity.TeamName) (entity.Teams, error)
	Change(ctx context.Context, t entity.Team) error
	Delete(ctx context.Context, teamId entity.Id) error
	Join(ctx context.Context, teamId entity.Id, userId entity.Id) error
}
