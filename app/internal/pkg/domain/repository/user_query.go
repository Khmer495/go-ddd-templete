package repository

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/ent"
	"github.com/Khmer495/go-templete/internal/pkg/ent/user"
)

func userQuerySelect(uq *ent.UserQuery) {
	uq.Select(
		user.FieldID,
		userSelector...,
	)
}

func userQueryFind(uq *ent.UserQuery, userId entity.Id) {
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

func userQueryPaging(uq *ent.UserQuery, limit entity.Limit, page entity.Page) {
	if limit != entity.NilLimit {
		uq.Limit(limit.Int())
		if page != entity.NilPage {
			uq.Offset((page.Int() - 1) * limit.Int())
		}
	}
}
