package handler

import (
	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
)

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
