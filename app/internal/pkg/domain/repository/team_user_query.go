package repository

import (
	"github.com/Khmer495/go-templete/internal/pkg/ent"
	"github.com/Khmer495/go-templete/internal/pkg/ent/teamuser"
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
