package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/user"
)

func userQuerySelect(uq *ent.UserQuery) {
	uq.Select(
		user.FieldID,
		userSelector...,
	)
}

func userQueryFind(uq *ent.UserQuery, userId model.Id) {
	uq.Where(
		uniqueUserPredicate(userId)...,
	)
}

func userQueryActive(uq *ent.UserQuery) {
	uq.Where(
		activeUserPredicate...,
	)
}

func userQueryOrder(uq *ent.UserQuery) {
	uq.Order(
		[]ent.OrderFunc{
			ent.Asc(user.FieldName),
		}...,
	)
}

func userQueryPaging(uq *ent.UserQuery, limit model.Limit, page model.Page) {
	if limit != model.NilLimit {
		uq.Limit(limit.Int())
		if page != model.NilPage {
			uq.Offset((page.Int() - 1) * limit.Int())
		}
	}
}
