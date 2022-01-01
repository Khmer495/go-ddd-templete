package testdata

import "github.com/Khmer495/go-templete/internal/pkg/domain/model"

var User0, _ = model.NewUser(Id0String, Datetime0Time, UserName0String)
var User0v2, _ = model.NewUser(Id0String, Datetime0Time, UserName1String)
var User1, _ = model.NewUser(Id1String, Datetime1Time, UserName1String)
var Users0 = model.Users{&User0}
var Users0_1 = model.Users{&User0, &User1}
