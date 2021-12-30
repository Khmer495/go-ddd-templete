package util

import (
	"bytes"
	"encoding/binary"
	"strings"
	"time"

	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

var JapaneseOneLetterRegexp = `a-zA-Z0-9０-９ａ-ｚＡ-Ｚ\p{Hiragana}\p{Katakana}ー、。\p{Han}`

var SymbolOnKeyBoardRegexp = "!\"#$%&'\\(\\)\\*\\+\\-\\.,\\/:;<=>\\?@\\[¥\\]\\^_`\\{\\|\\}~ ！”＃＄％＆’（）＊＋－．，／：；＜＝＞？＠￥＾＿～｛｜｝「」　"

func ContainsString(e string, s []string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func BytesToInt(val []byte) (int, error) {
	var i uint64
	buf := bytes.NewReader(val)
	if err := binary.Read(buf, binary.LittleEndian, &i); err != nil {
		return 0, cerror.Wrap(err, cerror.ErrorLevel, cerror.InterServerErrorCode, "binary.Read", "")
	}
	return int(i), nil
}

var JST *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)

func Truncate(t time.Time, d time.Duration) time.Time {
	tz := t.Location()
	tzDiff := time.Date(1, 0, 0, 0, 0, 0, 0, tz).Sub(time.Date(1, 0, 0, 0, 0, 0, 0, time.UTC))
	a := tzDiff.String()
	_ = a
	truncated := t.In(time.UTC).Add(-tzDiff).Truncate(d)
	return time.Date(truncated.Year(), truncated.Month(), truncated.Day(), truncated.Hour(), truncated.Minute(), truncated.Second(), truncated.Nanosecond(), tz)
}

func ToJSTDate(val time.Time) time.Time {
	return val.Add(time.Hour * 9).Truncate(time.Hour * 24)
}

func ToJSTDateInUTCTime(val time.Time) time.Time {
	return ToJSTDate(val).Add(-time.Hour * 9)
}

func isSomeEnv(env string, pattern []string) bool {
	return ContainsString(strings.ToLower(env), pattern)
}

func IsPrd(env string) bool {
	return isSomeEnv(env, []string{"prd", "prod", "production"})
}

func IsStg(env string) bool {
	return isSomeEnv(env, []string{"stg", "stage"})
}

func IsDev(env string) bool {
	return isSomeEnv(env, []string{"dev", "development"})
}

func IsLocal(env string) bool {
	return isSomeEnv(env, []string{"loc", "local"})
}
