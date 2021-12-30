package handler

import (
	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	GetProfile(ctx echo.Context) error
	PutProfile(ctx echo.Context) error
	PostUser(ctx echo.Context) error
}

type ITeamHandler interface {
	GetTeams(ctx echo.Context, params openapi.GetTeamsParams) error
	PostTeams(ctx echo.Context) error
	PutTeamsTeamId(ctx echo.Context, teamId string) error
	DeleteTeamsTeamId(ctx echo.Context, teamId string) error
	PostTeamsTeamIdJoin(ctx echo.Context, teamId string) error
}

type handler struct {
	IUserHandler
	ITeamHandler
}

func NewHandler(uh IUserHandler, th ITeamHandler) openapi.ServerInterface {
	return handler{
		IUserHandler: uh,
		ITeamHandler: th,
	}
}
