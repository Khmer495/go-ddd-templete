package model

import (
	"time"

	"github.com/Khmer495/go-templete/internal/pkg/util"
	"golang.org/x/xerrors"
)

type Datetime struct {
	time time.Time
}

var NilDatetime = Datetime{time.Time{}}

func NewDatetime(val time.Time) (Datetime, error) {
	// TODO: カスタムゼロ値は通さないようにしたい
	if val == NilDatetime.Time() {
		return NilDatetime, nil
	}
	return Datetime{val}, nil
}

var TimeNowFunc = time.Now

func InitDatetime() (Datetime, error) {
	datetime, err := NewDatetime(TimeNowFunc().UTC())
	if err != nil {
		return NilDatetime, xerrors.Errorf("NewDatetime: %w", err)
	}
	return datetime, nil
}

func (d Datetime) Time() time.Time {
	return d.time
}

func (d Datetime) PTime() *time.Time {
	if d == NilDatetime {
		return nil
	}
	dTime := d.Time()
	return &dTime
}

func (d Datetime) TimeStringISO8601() string {
	return d.Time().Format("2006-01-02T15:04:05-0700")
}

func (d Datetime) PTimeStringISO8601() *string {
	pTime := d.PTime()
	if pTime == nil {
		return nil
	}
	timeString := pTime.Format("2006-01-02T15:04:05-0700")
	return &timeString
}

func (d Datetime) Add(val time.Duration) Datetime {
	return Datetime{d.Time().Add(val)}
}

func (d Datetime) After(val Datetime) bool {
	return d.Time().After(val.Time())
}

func (d Datetime) Truncate(val time.Duration) Datetime {
	return Datetime{util.Truncate(d.Time(), val)}
}

func (d Datetime) Date() Datetime {
	return Datetime{util.Truncate(d.Time(), 24*time.Hour)}
}

func (d Datetime) UTC() Datetime {
	return Datetime{d.Time().In(time.UTC)}
}

func (d Datetime) JST() Datetime {
	return Datetime{d.Time().In(util.JST)}
}

func (d Datetime) JSTDateInUTCDatetime() Datetime {
	return d.JST().Date().UTC()
}
