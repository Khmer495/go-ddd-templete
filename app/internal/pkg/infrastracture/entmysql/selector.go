package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/team"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/teamuser"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/user"
)

var userSelector = []string{
	user.FieldUlid,
	user.FieldCreatedAt,
	user.FieldName,
}

var teamSelector = []string{
	team.FieldUlid,
	team.FieldCreateUserID,
	team.FieldName,
	team.FieldDescription,
}

var teamUserSelector = []string{
	teamuser.FieldTeamID,
	teamuser.FieldUserID,
}
