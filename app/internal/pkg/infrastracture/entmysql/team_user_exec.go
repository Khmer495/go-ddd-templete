package entmysql

import (
	"context"

	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/Khmer495/go-templete/internal/pkg/util/mysqlerror"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

func createTeamUser(ctx context.Context, tuc *ent.TeamUserClient, teamPk int, userPk int) (ent.TeamUser, error) {
	team_user, err := tuc.Create().
		SetTeamID(teamPk).
		SetUserID(userPk).
		Save(ctx)
	if err != nil {
		mysqlErr := &mysql.MySQLError{}
		if !xerrors.As(err, &mysqlErr) {
			return ent.TeamUser{}, cerror.WrapInternalServerError(err, "tuc.Create.Save")
		}
		if mysqlErr.Number == mysqlerror.DuplicateErrorNumber {
			return ent.TeamUser{}, cerror.Wrap(err, cerror.ErrorLevel, cerror.InvalidArgumentErrorCode, "mysqlErr.Number == mysqlerror.DuplicateErrorNumber", "既に加入しているチームです。")
		}
		return ent.TeamUser{}, cerror.WrapInternalServerError(err, "tuc.Create.Save")
	}
	return *team_user, nil
}
