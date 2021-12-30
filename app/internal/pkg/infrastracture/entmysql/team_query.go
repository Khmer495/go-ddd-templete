package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/team"
)

func teamQuerySelect(tq *ent.TeamQuery) {
	tq.Select(
		team.FieldID,
		teamSelector...,
	)
}

func teamQueryFind(tq *ent.TeamQuery, teamId entity.Id) {
	tq.Where(
		uniqueTeamPredicate(teamId)...,
	)
}

func teamQueryActive(tq *ent.TeamQuery) {
	tq.Where(
		activeTeamPredicate...,
	)
}

func teamQueryWith(tq *ent.TeamQuery) {
	tq.
		WithCreateUser(userQuerySelect).
		WithTeamUsers(teamUserQuerySelect, teamUserQueryWith)
}

func teamQueryOrder(tq *ent.TeamQuery) {
	tq.Order(
		[]ent.OrderFunc{
			ent.Asc(team.FieldID),
		}...,
	)
}

func teamQueryPaging(tq *ent.TeamQuery, limit entity.Limit, page entity.Page) {
	if limit != entity.NilLimit {
		tq.Limit(limit.Int())
		if page != entity.NilPage {
			tq.Offset((page.Int() - 1) * limit.Int())
		}
	}
}
