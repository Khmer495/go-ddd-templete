package testdata

import "github.com/Khmer495/go-templete/internal/pkg/domain/model"

var UserName0String = "aA0ａＡ０あアｱ阿"
var UserName0, _ = model.NewUserName(UserName0String)
var UserName1String = "bB1ｂＢ１いイｲ伊"
var UserName1, _ = model.NewUserName(UserName1String)
