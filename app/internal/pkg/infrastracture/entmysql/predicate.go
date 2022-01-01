package entmysql

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/predicate"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/team"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/user"
)

var uniqueUserPredicate = func(userId model.Id) []predicate.User {
	return []predicate.User{
		user.UlidEQ(userId.Ulid().String()),
	}
}

var activeUserPredicate = []predicate.User{
	user.DeletedAtIsNil(),
}

var uniqueTeamPredicate = func(teamId model.Id) []predicate.Team {
	return []predicate.Team{
		team.UlidEQ(teamId.Ulid().String()),
	}
}

var activeTeamPredicate = []predicate.Team{
	team.DeletedAtIsNil(),
}
