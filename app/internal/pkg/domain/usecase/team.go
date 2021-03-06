//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type ITeamUsecase interface {
	Create(ctx context.Context, name string, description string) (model.Team, error)
	GetList(ctx context.Context, limit int, page int, pName *string) (model.Teams, error)
	Change(ctx context.Context, teamdId string, pName *string, pDescription *string) (model.Team, error)
	Delete(ctx context.Context, teamId string) error
	Join(ctx context.Context, teamId string) error
}

type teamUsecase struct {
	ts service.ITeamService
	us service.IUserService
	tr repository.ITeamRepository
	ur repository.IUserRepository
}

func NewTeamUsecase(
	ts service.ITeamService,
	us service.IUserService,
	tr repository.ITeamRepository,
	ur repository.IUserRepository,
) ITeamUsecase {
	return teamUsecase{
		ts: ts,
		us: us,
		tr: tr,
		ur: ur,
	}
}

func (tu teamUsecase) Create(ctx context.Context, name string, description string) (model.Team, error) {
	createUser, err := tu.ur.Self(ctx)
	if err != nil {
		return model.NilTeam, cerror.WrapInternalServerError(err, "tu.ur.GetSelf")
	}

	team, err := model.InitTeam(createUser, name, description)
	if err != nil {
		return model.NilTeam, parsemodelConstractorError(err, "model.InitTeam")
	}

	err = tu.tr.Create(ctx, team)
	if err != nil {
		return model.NilTeam, cerror.WrapInternalServerError(err, "tu.tr.Create")
	}

	return team, nil
}

func (tu teamUsecase) GetList(ctx context.Context, limitInt int, pageInt int, pNameString *string) (model.Teams, error) {
	limit, err := model.NewLimit(limitInt)
	if err != nil {
		return model.NilTeams, parsemodelConstractorError(err, "model.NewLimit")
	}
	page, err := model.NewPage(pageInt)
	if err != nil {
		return model.NilTeams, parsemodelConstractorError(err, "model.NewPage")
	}

	if pNameString == nil {
		teams, err := tu.tr.List(ctx, limit, page)
		if err != nil {
			return model.NilTeams, cerror.WrapInternalServerError(err, "tu.tr.List")
		}
		return teams, nil
	} else {
		name, err := model.NewTeamName(*pNameString)
		if err != nil {
			return model.NilTeams, parsemodelConstractorError(err, "model.NewTeamName")
		}

		teams, err := tu.tr.SearchByNamePrefix(ctx, limit, page, name)
		if err != nil {
			return model.NilTeams, cerror.WrapInternalServerError(err, "tu.tr.SearchByNamePrefix")
		}

		return teams, nil
	}
}

func (tu teamUsecase) Change(ctx context.Context, teamIdString string, pName *string, pDescription *string) (model.Team, error) {
	teamId, err := model.NewId(teamIdString)
	if err != nil {
		return model.NilTeam, cerror.WrapNotFoundError(err, "model.NewId", "??????????????????????????????????????????")
	}
	team, err := tu.tr.One(ctx, teamId)
	if err != nil {
		return model.NilTeam, cerror.WrapNotFoundError(err, "tu.tr.One", "??????????????????????????????????????????")
	}

	if pName != nil {
		err := team.SetName(*pName)
		if err != nil {
			return model.NilTeam, parsemodelConstractorError(err, "team.SetName")
		}
	}
	if pDescription != nil {
		err := team.SetDesctiprion(*pDescription)
		if err != nil {
			return model.NilTeam, parsemodelConstractorError(err, "team.SetDesctiprion")
		}
	}

	err = tu.tr.Change(ctx, team)
	if err != nil {
		return model.NilTeam, cerror.WrapInternalServerError(err, "tu.tr.Change")
	}

	return team, nil
}

func (tu teamUsecase) Delete(ctx context.Context, teamIdString string) error {
	teamId, err := model.NewId(teamIdString)
	if err != nil {
		return cerror.WrapNotFoundError(err, "model.NewId", "??????????????????????????????????????????")
	}
	team, err := tu.tr.One(ctx, teamId)
	if err != nil {
		return cerror.WrapNotFoundError(err, "tu.tr.One", "??????????????????????????????????????????")
	}

	err = tu.tr.Delete(ctx, team.Id())
	if err != nil {
		return cerror.WrapInternalServerError(err, "tu.tr.Delete")
	}

	return nil
}

func (tu teamUsecase) Join(ctx context.Context, teamIdString string) error {
	teamId, err := model.NewId(teamIdString)
	if err != nil {
		return cerror.WrapNotFoundError(err, "model.NewId", "??????????????????????????????????????????")
	}
	team, err := tu.tr.One(ctx, teamId)
	if err != nil {
		return cerror.WrapNotFoundError(err, "tu.tr.One", "??????????????????????????????????????????")
	}

	userId, err := tu.us.SelfId(ctx)
	if err != nil {
		return cerror.WrapInternalServerError(err, "tu.ur.Self")
	}

	isTeamUser := team.IsUser(userId)
	if isTeamUser {
		return cerror.NewAlreadyExistError("team.IsUser", "??????????????????????????????????????????")
	}

	err = tu.tr.Join(ctx, teamId, userId)
	if err != nil {
		return cerror.WrapInternalServerError(err, "tu.tr.Join")
	}

	return nil
}
