package testdata

import "github.com/Khmer495/go-templete/internal/pkg/domain/entity"

var Team0, _ = entity.NewTeam(Id0String, Datetime0Time, User0, TeamName0String, TeamDescription0String, Users0_1)
var Team0v2, _ = entity.NewTeam(Id0String, Datetime0Time, User0, TeamName1String, TeamDescription1String, Users0_1)
var Team0v3, _ = entity.NewTeam(Id0String, Datetime0Time, User0, TeamName0String, TeamDescription0String, Users0)
var Teams0 = entity.Teams{&Team0}
