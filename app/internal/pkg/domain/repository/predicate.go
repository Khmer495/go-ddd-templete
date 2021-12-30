package repository

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/ent/predicate"
	"github.com/Khmer495/go-templete/internal/pkg/ent/team"
	"github.com/Khmer495/go-templete/internal/pkg/ent/user"
)

var uniqueUserPredicate = func(userId entity.Id) []predicate.User {
	return []predicate.User{
		user.UlidEQ(userId.Ulid().String()),
	}
}

var activeUserPredicate = []predicate.User{
	user.DeletedAtIsNil(),
}

var uniqueTeamPredicate = func(teamId entity.Id) []predicate.Team {
	return []predicate.Team{
		team.UlidEQ(teamId.Ulid().String()),
	}
}

var activeTeamPredicate = []predicate.Team{
	team.DeletedAtIsNil(),
}
