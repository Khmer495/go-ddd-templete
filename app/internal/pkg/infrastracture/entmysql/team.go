package entmysql

import (
	"context"

	entity "github.com/Khmer495/go-templete/internal/pkg/domain/entity"
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

func (tr teamRepository) Create(ctx context.Context, t entity.Team) error {
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

func (tr teamRepository) IsExist(ctx context.Context, teamId entity.Id) (bool, error) {
	is, err := isTeamExist(ctx, tr.mysqlClient.Team, teamId)
	if err != nil {
		return false, xerrors.Errorf("isTeamExist: %w", err)
	}
	return is, nil
}

func (tr teamRepository) One(ctx context.Context, teamId entity.Id) (entity.Team, error) {
	team, err := findTeam(ctx, tr.mysqlClient.Team, teamId)
	if err != nil {
		return entity.NilTeam, xerrors.Errorf("findTeam: %w", err)
	}
	entityTeam, err := team.EntityTeam()
	if err != nil {
		return entity.NilTeam, xerrors.Errorf("team.EntityTeam: %w", err)
	}
	return entityTeam, nil
}

func (tr teamRepository) List(ctx context.Context, limit entity.Limit, page entity.Page) (entity.Teams, error) {
	teams, err := getTeams(ctx, tr.mysqlClient.Team, limit, page, getTeamsParam{})
	if err != nil {
		return entity.NilTeams, xerrors.Errorf("getTeams: %w", err)
	}
	entityTeams, err := ent.Teams(teams).EntityTeams()
	if err != nil {
		return entity.NilTeams, xerrors.Errorf("teams.NewEntityTeams: %w", err)
	}
	return entityTeams, nil
}

func (tr teamRepository) SearchByNamePrefix(ctx context.Context, limit entity.Limit, page entity.Page, tn entity.TeamName) (entity.Teams, error) {
	teams, err := getTeams(ctx, tr.mysqlClient.Team, limit, page, getTeamsParam{pTeamNamePrefix: &tn})
	if err != nil {
		return entity.NilTeams, xerrors.Errorf("getTeams: %w", err)
	}
	entityTeams, err := ent.Teams(teams).EntityTeams()
	if err != nil {
		return entity.NilTeams, xerrors.Errorf("teams.NewEntityTeams: %w", err)
	}
	return entityTeams, nil
}

func (tr teamRepository) Change(ctx context.Context, t entity.Team) error {
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

func (tr teamRepository) Delete(ctx context.Context, teamId entity.Id) error {
	err := func() error {
		tx, err := tr.mysqlClient.Tx(ctx)
		if err != nil {
			return cerror.WrapInternalServerError(err, "tr.mysqlClient.Tx")
		}
		team, err := findTeamRecord(ctx, tx.Team, teamId)
		if err != nil {
			return txRollbackAndParseError(tx, err, "findTeamRecord")
		}
		deletedAt, err := entity.InitDatetime()
		if err != nil {
			return txRollbackAndParseError(tx, err, "entity.InitDatetime")
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

func (tr teamRepository) Join(ctx context.Context, teamId entity.Id, userId entity.Id) error {
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
