//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type ITeamUsecase interface {
	Create(ctx context.Context, name string, description string) (entity.Team, error)
	GetList(ctx context.Context, limit int, page int, pName *string) (entity.Teams, error)
	Change(ctx context.Context, teamdId string, pName *string, pDescription *string) (entity.Team, error)
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

func (tu teamUsecase) Create(ctx context.Context, name string, description string) (entity.Team, error) {
	createUser, err := tu.ur.Self(ctx)
	if err != nil {
		return entity.NilTeam, cerror.WrapInternalServerError(err, "tu.ur.GetSelf")
	}

	team, err := entity.InitTeam(createUser, name, description)
	if err != nil {
		return entity.NilTeam, parseEntityConstractorError(err, "entity.InitTeam")
	}

	err = tu.tr.Create(ctx, team)
	if err != nil {
		return entity.NilTeam, cerror.WrapInternalServerError(err, "tu.tr.Create")
	}

	return team, nil
}

func (tu teamUsecase) GetList(ctx context.Context, limitInt int, pageInt int, pNameString *string) (entity.Teams, error) {
	limit, err := entity.NewLimit(limitInt)
	if err != nil {
		return entity.NilTeams, parseEntityConstractorError(err, "entity.NewLimit")
	}
	page, err := entity.NewPage(pageInt)
	if err != nil {
		return entity.NilTeams, parseEntityConstractorError(err, "entity.NewPage")
	}

	if pNameString == nil {
		teams, err := tu.tr.List(ctx, limit, page)
		if err != nil {
			return entity.NilTeams, cerror.WrapInternalServerError(err, "tu.tr.List")
		}
		return teams, nil
	} else {
		name, err := entity.NewTeamName(*pNameString)
		if err != nil {
			return entity.NilTeams, parseEntityConstractorError(err, "entity.NewTeamName")
		}

		teams, err := tu.tr.SearchByNamePrefix(ctx, limit, page, name)
		if err != nil {
			return entity.NilTeams, cerror.WrapInternalServerError(err, "tu.tr.SearchByNamePrefix")
		}

		return teams, nil
	}
}

func (tu teamUsecase) Change(ctx context.Context, teamIdString string, pName *string, pDescription *string) (entity.Team, error) {
	teamId, err := entity.NewId(teamIdString)
	if err != nil {
		return entity.NilTeam, cerror.WrapNotFoundError(err, "entity.NewId", "指定のチームが存在しません。")
	}
	team, err := tu.tr.One(ctx, teamId)
	if err != nil {
		return entity.NilTeam, cerror.WrapNotFoundError(err, "tu.tr.One", "指定のチームが存在しません。")
	}

	if pName != nil {
		err := team.SetName(*pName)
		if err != nil {
			return entity.NilTeam, parseEntityConstractorError(err, "team.SetName")
		}
	}
	if pDescription != nil {
		err := team.SetDesctiprion(*pDescription)
		if err != nil {
			return entity.NilTeam, parseEntityConstractorError(err, "team.SetDesctiprion")
		}
	}

	err = tu.tr.Change(ctx, team)
	if err != nil {
		return entity.NilTeam, cerror.WrapInternalServerError(err, "tu.tr.Change")
	}

	return team, nil
}

func (tu teamUsecase) Delete(ctx context.Context, teamIdString string) error {
	teamId, err := entity.NewId(teamIdString)
	if err != nil {
		return cerror.WrapNotFoundError(err, "entity.NewId", "指定のチームが存在しません。")
	}
	team, err := tu.tr.One(ctx, teamId)
	if err != nil {
		return cerror.WrapNotFoundError(err, "tu.tr.One", "指定のチームが存在しません。")
	}

	err = tu.tr.Delete(ctx, team.Id())
	if err != nil {
		return cerror.WrapInternalServerError(err, "tu.tr.Delete")
	}

	return nil
}

func (tu teamUsecase) Join(ctx context.Context, teamIdString string) error {
	teamId, err := entity.NewId(teamIdString)
	if err != nil {
		return cerror.WrapNotFoundError(err, "entity.NewId", "指定のチームが存在しません。")
	}
	team, err := tu.tr.One(ctx, teamId)
	if err != nil {
		return cerror.WrapNotFoundError(err, "tu.tr.One", "指定のチームが存在しません。")
	}

	userId, err := tu.us.SelfId(ctx)
	if err != nil {
		return cerror.WrapInternalServerError(err, "tu.ur.Self")
	}

	isTeamUser := team.IsUser(userId)
	if isTeamUser {
		return cerror.NewAlreadyExistError("team.IsUser", "既にチームに参加しています。")
	}

	err = tu.tr.Join(ctx, teamId, userId)
	if err != nil {
		return cerror.WrapInternalServerError(err, "tu.tr.Join")
	}

	return nil
}
