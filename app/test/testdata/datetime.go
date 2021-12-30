package testdata

import (
	"time"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/util"
)

func init() {
	entity.TimeNowFunc = func() time.Time {
		return Datetime0Time
	}
}

var Datetime0Time time.Time = time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
var Datetime0StringISO8601 = "2006-01-02T15:04:05+0000"
var Datetime0, _ = entity.NewDatetime(Datetime0Time)
var JstDatetime0Time time.Time = time.Date(2006, 1, 2, 15, 4, 5, 0, util.JST)
var JstDatetime0StringISO8601 = "2006-01-03T00:04:05+0900"
var JstDatetime0, _ = entity.NewDatetime(JstDatetime0Time)
var Datetime1Time time.Time = time.Date(2020, 4, 1, 12, 34, 56, 0, time.UTC)
var Datetime1, _ = entity.NewDatetime(Datetime1Time)

// Datetime0の1年前
var Datetime2Time time.Time = time.Date(2005, 1, 2, 15, 4, 5, 0, time.UTC)
var Datetime2, _ = entity.NewDatetime(Datetime2Time)
