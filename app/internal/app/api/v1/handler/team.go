package handler

import (
	"context"
	"net/http"

	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/usecase"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type teamHandler struct {
	tu usecase.ITeamUsecase
}

func NewTeamHandler(tu usecase.ITeamUsecase) ITeamHandler {
	return teamHandler{
		tu: tu,
	}
}

func fromEntityTeamToOpenapiTeam(ctx context.Context, t entity.Team) (openapi.Team, error) {
	users, err := fromEntityUsersToOpenapiUsers(ctx, t.Users())
	if err != nil {
		return openapi.Team{}, xerrors.Errorf("fromEntityUsersToOpenapiUsers: %w", err)
	}
	return openapi.Team{
		Id:          t.Id().Ulid().String(),
		Name:        t.Name().String(),
		Description: t.Description().String(),
		Users:       users,
	}, nil
}

func fromEntityTeamsToOpenapiTeams(ctx context.Context, ts entity.Teams) ([]openapi.Team, error) {
	teams := []openapi.Team{}
	for _, t := range ts {
		team, err := fromEntityTeamToOpenapiTeam(ctx, *t)
		if err != nil {
			return nil, xerrors.Errorf("fromEntityTeamToOpenapiTeam: %w", err)
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (th teamHandler) GetTeams(ctx echo.Context, params openapi.GetTeamsParams) error {
	teams, err := th.tu.GetList(ctx.Request().Context(), params.Limit, params.Page, params.Name)
	if err != nil {
		zap.S().Errorf("th.tu.GetList: %+v", err)
		if cerror.IsCode(err, cerror.InvalidArgumentErrorCode) {
			return NewInvalidArgumentError(ctx, cerror.As(err).ClientMsg())
		}
		return NewInternalServerError(ctx)
	}
	res, err := fromEntityTeamsToOpenapiTeams(ctx.Request().Context(), teams)
	if err != nil {
		zap.S().Errorf("fromEntityTeamsToOpenapiTeams: %+v", err)
		return NewInternalServerError(ctx)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (th teamHandler) PostTeams(ctx echo.Context) error {
	req := openapi.PostTeamsJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		zap.S().Errorf("ctx.Bind: %+v", err)
		return NewInvalidRequestFormatError(ctx, err)
	}
	team, err := th.tu.Create(ctx.Request().Context(), req.Name, req.Description)
	if err != nil {
		zap.S().Errorf("th.tu.Create: %+v", err)
		if cerror.IsCode(err, cerror.InvalidArgumentErrorCode) {
			return NewInvalidArgumentError(ctx, cerror.As(err).ClientMsg())
		}
		return NewInternalServerError(ctx)
	}
	res, err := fromEntityTeamToOpenapiTeam(ctx.Request().Context(), team)
	if err != nil {
		zap.S().Errorf("fromEntityTeamToOpenapiTeam: %+v", err)
		return NewInternalServerError(ctx)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (th teamHandler) PutTeamsTeamId(ctx echo.Context, teamId string) error {
	req := openapi.PutTeamsTeamIdJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		zap.S().Errorf("ctx.Bind: %+v", err)
		return NewInvalidRequestFormatError(ctx, err)
	}

	team, err := th.tu.Change(ctx.Request().Context(), teamId, req.Name, req.Description)
	if err != nil {
		zap.S().Errorf("th.tu.Create: %+v", err)
		if cerror.IsCode(err, cerror.InvalidArgumentErrorCode) {
			return NewInvalidArgumentError(ctx, cerror.As(err).ClientMsg())
		}
		return NewInternalServerError(ctx)
	}
	res, err := fromEntityTeamToOpenapiTeam(ctx.Request().Context(), team)
	if err != nil {
		zap.S().Errorf("fromEntityTeamToOpenapiTeam: %+v", err)
		return NewInternalServerError(ctx)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (th teamHandler) DeleteTeamsTeamId(ctx echo.Context, teamId string) error {
	err := th.tu.Delete(ctx.Request().Context(), teamId)
	if err != nil {
		zap.S().Errorf("th.tu.Delete: %+v", err)
		if cerror.IsCode(err, cerror.UnauthorizedErrorCode) {
			return NewNotFoundError(ctx)
		}
		if cerror.IsCode(err, cerror.NotFoundErrorCode) {
			return NewNotFoundError(ctx)
		}
		return NewInternalServerError(ctx)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (th teamHandler) PostTeamsTeamIdJoin(ctx echo.Context, teamId string) error {
	err := th.tu.Join(ctx.Request().Context(), teamId)
	if err != nil {
		zap.S().Errorf("th.tu.Join: %+v", err)
		if cerror.IsCode(err, cerror.AlreadyExistsErrorCode) {
			return NewAlreadyAppliedError(ctx)
		}
		if cerror.IsCode(err, cerror.UnauthorizedErrorCode) {
			return NewNotFoundError(ctx)
		}
		if cerror.IsCode(err, cerror.NotFoundErrorCode) {
			return NewNotFoundError(ctx)
		}
		return NewInternalServerError(ctx)
	}
	return ctx.NoContent(http.StatusNoContent)
}
