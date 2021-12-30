package testdata

import "github.com/Khmer495/go-templete/internal/pkg/domain/entity"

var TeamName0String = "aA0ａＡ０あアｱ阿! ！　"
var TeamName0, _ = entity.NewTeamName(TeamName0String)
var TeamName1String = "bB1ｂＢ１いイｲ伊! ！　"
var TeamName1, _ = entity.NewTeamName(TeamName1String)
