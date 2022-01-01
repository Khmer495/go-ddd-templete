package model

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Khmer495/go-templete/internal/pkg/util"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type UserName struct {
	string
}

var NilUserName = UserName{""}

func NewUserName(val string) (UserName, error) {
	val = strings.Replace(val, "　", " ", -1)                                //全角スペースを半角スペースに置換
	val = strings.TrimSpace(val)                                            //前後のスペース削除
	regexpPattern := `^[` + util.JapaneseOneLetterRegexp + ` ` + `]{1,16}$` // 半角スペースも含む
	isMatch, err := regexp.MatchString(regexpPattern, val)
	if err != nil {
		return NilUserName, cerror.WrapInternalServerError(err, "regexp.MatchString")
	}
	if !isMatch {
		return NilUserName, cerror.NewInvalidArgumentError("regexp.MatchString", fmt.Sprintf("ユーザー名は1文字以上16文字以下の半角全角英数字ひらがなカタカナ漢字及び半角スペースで構成されます。\nwant: %shave: %s", regexpPattern, val))
	}
	return UserName{val}, nil
}

func (un UserName) String() string {
	return un.string
}
