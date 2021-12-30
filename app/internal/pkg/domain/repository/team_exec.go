package repository

import (
	"context"
	"fmt"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/ent"
	"github.com/Khmer495/go-templete/internal/pkg/ent/team"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/Khmer495/go-templete/internal/pkg/util/mysqlerror"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

func createTeam(ctx context.Context, tc *ent.TeamClient, t entity.Team, createUserPk int) (ent.Team, error) {
	team, err := tc.Create().
		SetUlid(t.Id().Ulid().String()).
		SetCreatedAt(t.CreatedAt().Time()).
		SetUpdatedAt(t.CreatedAt().Time()).
		SetCreateUserID(createUserPk).
		SetName(t.Name().String()).
		SetDescription(t.Description().String()).
		Save(ctx)
	if err != nil {
		mysqlErr := &mysql.MySQLError{}
		if !xerrors.As(err, &mysqlErr) {
			return ent.Team{}, cerror.WrapInternalServerError(err, "tc.Create.Save")
		}
		if mysqlErr.Number == mysqlerror.DuplicateErrorNumber {
			return ent.Team{}, cerror.Wrap(err, cerror.ErrorLevel, cerror.InvalidArgumentErrorCode, "mysqlErr.Number == mysqlerror.DuplicateErrorNumber", "既に登録されているチームです。")
		}
		return ent.Team{}, cerror.WrapInternalServerError(err, "tc.Create.Save")
	}
	return *team, nil
}

func isTeamExist(ctx context.Context, tc *ent.TeamClient, teamId entity.Id) (bool, error) {
	tq := tc.Query()
	teamQueryFind(tq, teamId)
	teamQueryActive(tq)
	ok, err := tq.Exist(ctx)
	if err != nil {
		return false, cerror.WrapInternalServerError(err, "tq.Exist")
	}
	return ok, nil
}

func findTeamRecord(ctx context.Context, tc *ent.TeamClient, teamId entity.Id) (ent.Team, error) {
	tq := tc.Query()
	teamQuerySelect(tq)
	teamQueryFind(tq, teamId)
	teamQueryActive(tq)
	team, err := tq.Only(ctx)
	if err != nil {
		return ent.Team{}, parseQueryOnlyByUlid(err, "tq.Only", fmt.Sprintf("teamId: %sが存在しません。", teamId.Ulid().String()))
	}
	return *team, nil
}

func findTeam(ctx context.Context, tc *ent.TeamClient, teamId entity.Id) (ent.Team, error) {
	tq := tc.Query()
	teamQuerySelect(tq)
	teamQueryFind(tq, teamId)
	teamQueryActive(tq)
	teamQueryWith(tq)
	team, err := tq.Only(ctx)
	if err != nil {
		return ent.Team{}, parseQueryOnlyByUlid(err, "tq.Only", fmt.Sprintf("teamId: %sが存在しません。", teamId.Ulid().String()))
	}
	return *team, nil
}

type getTeamsParam struct {
	pTeamNamePrefix *entity.TeamName
}

func getTeams(ctx context.Context, tc *ent.TeamClient, limit entity.Limit, page entity.Page, gtp getTeamsParam) (ent.Teams, error) {
	tq := tc.Query()
	teamQuerySelect(tq)
	teamQueryActive(tq)
	teamQueryOrder(tq)
	teamQueryWith(tq)
	teamQueryPaging(tq, limit, page)
	if gtp.pTeamNamePrefix != nil {
		tq.Where(
			team.NameHasPrefix(gtp.pTeamNamePrefix.String()),
		)
	}
	teams, err := tq.All(ctx)
	if err != nil {
		return nil, cerror.WrapInternalServerError(err, "tq.All")
	}
	return teams, nil
}

func updateTeam(ctx context.Context, tc *ent.TeamClient, teamPk int, t entity.Team) (ent.Team, error) {
	team, err := tc.UpdateOneID(teamPk).
		SetName(t.Name().String()).
		SetDescription(t.Description().String()).
		Save(ctx)
	if err != nil {
		return ent.Team{}, cerror.WrapInternalServerError(err, "tc.UpdateOneID.Save")
	}
	return *team, nil
}

func deleteTeam(ctx context.Context, tc *ent.TeamClient, teamPk int, deletedAt entity.Datetime) error {
	_, err := tc.UpdateOneID(teamPk).
		SetDeletedAt(deletedAt.Time()).
		Save(ctx)
	if err != nil {
		return cerror.WrapInternalServerError(err, "mc.DeleteOneID.Exec")
	}
	return nil
}
