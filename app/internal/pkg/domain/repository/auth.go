//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
)

type IAuthRepository interface {
	FindUserPkByFirebaseUserId(ctx context.Context) (int, error)
	FindUserIdByFirebaseUserId(ctx context.Context) (entity.Id, error)
}
