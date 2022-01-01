package handler

import (
	"context"

	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"golang.org/x/xerrors"
)

func frommodelUserToOpenapiUser(ctx context.Context, u model.User) (openapi.User, error) {
	return openapi.User{
		Id:   u.Id().Ulid().String(),
		Name: u.Name().String(),
	}, nil
}

func frommodelUsersToOpenapiUsers(ctx context.Context, us model.Users) ([]openapi.User, error) {
	users := []openapi.User{}
	for _, u := range us {
		user, err := frommodelUserToOpenapiUser(ctx, *u)
		if err != nil {
			return nil, xerrors.Errorf("frommodelUserToOpenapiUser: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}
