package entmysql

import (
	"context"

	model "github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"golang.org/x/xerrors"
)

type teamRepository struct {
	mysqlClient *ent.Client
}

func NewTeamRepository() repository.ITeamRepository {
	return teamRepository{
		mysqlClient: MysqlClient,
	}
}

func (tr teamRepository) Create(ctx context.Context, t model.Team) error {
	err := func() error {
		tx, err := tr.mysqlClient.Tx(ctx)
		if err != nil {
			return xerrors.Errorf("tr.mysqlClient.Tx: %w", err)
		}
		createUser, err := findUser(ctx, tx.User, t.CreateUser().Id())
		if err != nil {
			return txRollbackAndParseError(tx, err, "findUser")
		}
		team, err := createTeam(ctx, tx.Team, t, createUser.ID)
		if err != nil {
			return txRollbackAndParseError(tx, err, "createTeam")
		}
		_, err = createTeamUser(ctx, tx.TeamUser, team.ID, createUser.ID)
		if err != nil {
			return txRollbackAndParseError(tx, err, "createTeamUser")
		}
		return tx.Commit()
	}()
	if err != nil {
		return newTxCommitError(err)
	}
	return nil
}

func (tr teamRepository) IsExist(ctx context.Context, teamId model.Id) (bool, error) {
	is, err := isTeamExist(ctx, tr.mysqlClient.Team, teamId)
	if err != nil {
		return false, xerrors.Errorf("isTeamExist: %w", err)
	}
	return is, nil
}

func (tr teamRepository) One(ctx context.Context, teamId model.Id) (model.Team, error) {
	team, err := findTeam(ctx, tr.mysqlClient.Team, teamId)
	if err != nil {
		return model.NilTeam, xerrors.Errorf("findTeam: %w", err)
	}
	modelTeam, err := entTeamTomodelTeam(team)
	if err != nil {
		return model.NilTeam, xerrors.Errorf("team.modelTeam: %w", err)
	}
	return modelTeam, nil
}

func (tr teamRepository) List(ctx context.Context, limit model.Limit, page model.Page) (model.Teams, error) {
	teams, err := getTeams(ctx, tr.mysqlClient.Team, limit, page, getTeamsParam{})
	if err != nil {
		return model.NilTeams, xerrors.Errorf("getTeams: %w", err)
	}
	modelTeams, err := entTeamsTomodelTeams(teams)
	if err != nil {
		return model.NilTeams, xerrors.Errorf("teams.NewmodelTeams: %w", err)
	}
	return modelTeams, nil
}

func (tr teamRepository) SearchByNamePrefix(ctx context.Context, limit model.Limit, page model.Page, tn model.TeamName) (model.Teams, error) {
	teams, err := getTeams(ctx, tr.mysqlClient.Team, limit, page, getTeamsParam{pTeamNamePrefix: &tn})
	if err != nil {
		return model.NilTeams, xerrors.Errorf("getTeams: %w", err)
	}
	modelTeams, err := entTeamsTomodelTeams(teams)
	if err != nil {
		return model.NilTeams, xerrors.Errorf("teams.NewmodelTeams: %w", err)
	}
	return modelTeams, nil
}

func (tr teamRepository) Change(ctx context.Context, t model.Team) error {
	err := func() error {
		tx, err := tr.mysqlClient.Tx(ctx)
		if err != nil {
			return xerrors.Errorf("tr.mysqlClient.Tx: %w", err)
		}
		team, err := findTeam(ctx, tx.Team, t.Id())
		if err != nil {
			return txRollbackAndParseError(tx, err, "findTeam")
		}
		_, err = updateTeam(ctx, tx.Team, team.ID, t)
		if err != nil {
			return txRollbackAndParseError(tx, err, "updateTeam")
		}
		return tx.Commit()
	}()
	if err != nil {
		return newTxCommitError(err)
	}
	return nil
}

func (tr teamRepository) Delete(ctx context.Context, teamId model.Id) error {
	err := func() error {
		tx, err := tr.mysqlClient.Tx(ctx)
		if err != nil {
			return cerror.WrapInternalServerError(err, "tr.mysqlClient.Tx")
		}
		team, err := findTeamRecord(ctx, tx.Team, teamId)
		if err != nil {
			return txRollbackAndParseError(tx, err, "findTeamRecord")
		}
		deletedAt, err := model.InitDatetime()
		if err != nil {
			return txRollbackAndParseError(tx, err, "model.InitDatetime")
		}
		err = deleteTeam(ctx, tx.Team, team.ID, deletedAt)
		if err != nil {
			return txRollbackAndParseError(tx, err, "deleteTeam")
		}
		return tx.Commit()
	}()
	if err != nil {
		return newTxCommitError(err)
	}
	return nil
}

func (tr teamRepository) Join(ctx context.Context, teamId model.Id, userId model.Id) error {
	err := func() error {
		tx, err := tr.mysqlClient.Tx(ctx)
		if err != nil {
			return xerrors.Errorf("tr.mysqlClient.Tx: %w", err)
		}
		team, err := findTeamRecord(ctx, tx.Team, teamId)
		if err != nil {
			return txRollbackAndParseError(tx, err, "findTeamRecord")
		}
		user, err := findUserRecord(ctx, tx.User, userId)
		if err != nil {
			return txRollbackAndParseError(tx, err, "findUserRecord")
		}
		_, err = createTeamUser(ctx, tx.TeamUser, team.ID, user.ID)
		if err != nil {
			return txRollbackAndParseError(tx, err, "createTeamUser")
		}
		return tx.Commit()
	}()
	if err != nil {
		return newTxCommitError(err)
	}
	return nil
}
