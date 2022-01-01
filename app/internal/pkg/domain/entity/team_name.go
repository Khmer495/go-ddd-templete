package model

import (
	"fmt"
	"regexp"

	"github.com/Khmer495/go-templete/internal/pkg/util"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type TeamName struct {
	string
}

var NilTeamName = TeamName{""}

func NewTeamName(val string) (TeamName, error) {
	regexpPattern := `^[` + util.JapaneseOneLetterRegexp + util.SymbolOnKeyBoardRegexp + `]{1,16}$`
	isMatch, err := regexp.MatchString(regexpPattern, val)
	if err != nil {
		return NilTeamName, cerror.WrapInternalServerError(err, "regexp.MatchString")
	}
	if !isMatch {
		return NilTeamName, cerror.NewInvalidArgumentError("regexp.MatchString", fmt.Sprintf("チーム名は1文字以上16文字以下の半角全角英数字ひらがなカタカナ漢字で構成されます。\nwant: %s\nhave: %s", regexpPattern, val))
	}
	return TeamName{val}, nil
}

func (tn TeamName) String() string {
	return tn.string
}
