package entity

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Khmer495/go-templete/internal/pkg/util"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type TeamDescription struct {
	string
}

var NilTeamDescription = TeamDescription{strings.Repeat("1234567890", 25) + "1234567"}

func NewTeamDescription(val string) (TeamDescription, error) {
	regexpPattern := `^[` + util.JapaneseOneLetterRegexp + util.SymbolOnKeyBoardRegexp + `\n]{0,256}$`
	isMatch, err := regexp.MatchString(regexpPattern, val)
	if err != nil {
		return NilTeamDescription, cerror.WrapInternalServerError(err, "regexp.MatchString")
	}
	if !isMatch {
		return NilTeamDescription, cerror.NewInvalidArgumentError("regexp.MatchString", fmt.Sprintf("チーム説明は1文字以上256文字以下の半角全角英数字ひらがなカタカナ漢字で構成されます。\nwant: %s\nhave: %s", regexpPattern, val))
	}
	return TeamDescription{val}, nil
}

func (td TeamDescription) String() string {
	return td.string
}
