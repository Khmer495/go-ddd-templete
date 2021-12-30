//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
)

type IUserRepository interface {
	Register(ctx context.Context, u entity.User) error
	IsExist(ctx context.Context, userId entity.Id) (bool, error)
	Self(ctx context.Context) (entity.User, error)
	List(ctx context.Context, limit entity.Limit, page entity.Page) (entity.Users, error)
	Select(ctx context.Context, limit entity.Limit, page entity.Page, ids entity.Ids) (entity.Users, error)
	SearchByNamePrefix(ctx context.Context, limit entity.Limit, page entity.Page, un entity.UserName) (entity.Users, error)
	Change(ctx context.Context, u entity.User) error
}
