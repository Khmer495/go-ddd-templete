//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
)

type IUserRepository interface {
	Register(ctx context.Context, u model.User) error
	IsExist(ctx context.Context, userId model.Id) (bool, error)
	Self(ctx context.Context) (model.User, error)
	List(ctx context.Context, limit model.Limit, page model.Page) (model.Users, error)
	Select(ctx context.Context, limit model.Limit, page model.Page, ids model.Ids) (model.Users, error)
	SearchByNamePrefix(ctx context.Context, limit model.Limit, page model.Page, un model.UserName) (model.Users, error)
	Change(ctx context.Context, u model.User) error
}
