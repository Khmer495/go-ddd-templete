package testdata

import "github.com/Khmer495/go-templete/internal/pkg/domain/entity"

var User0, _ = entity.NewUser(Id0String, Datetime0Time, UserName0String)
var User0v2, _ = entity.NewUser(Id0String, Datetime0Time, UserName1String)
var User1, _ = entity.NewUser(Id1String, Datetime1Time, UserName1String)
var Users0 = entity.Users{&User0}
var Users0_1 = entity.Users{&User0, &User1}
