package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/teamuser"
)

func teamUserQuerySelect(tuq *ent.TeamUserQuery) {
	tuq.Select(
		teamuser.FieldID,
		teamUserSelector...,
	)
}

func teamUserQueryWith(tuq *ent.TeamUserQuery) {
	tuq.
		WithUser(userQuerySelect, userQueryActive, userQueryOrder)
}
