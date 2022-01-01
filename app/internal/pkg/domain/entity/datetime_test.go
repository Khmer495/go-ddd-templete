package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/Khmer495/go-templete/internal/pkg/util"
)

var datetime0Time time.Time = time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
var datetime0StringISO8601 = "2006-01-02T15:04:05+0000"
var datetime0 = Datetime{datetime0Time}
var jstDatetime0Time time.Time = time.Date(2006, 1, 3, 0, 4, 5, 0, time.UTC)
var jstDatetime0StringISO8601 = "2006-01-03T00:04:05+0000"
var jstDatetime0 = Datetime{jstDatetime0Time}
var datetime1Time time.Time = time.Date(2020, 4, 1, 12, 34, 56, 0, time.UTC)
var datetime1 = Datetime{datetime1Time}

func init() {
	TimeNowFunc = func() time.Time {
		return datetime0Time
	}
}

func TestNewDatetime(t *testing.T) {
	type args struct {
		val time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    Datetime
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: datetime0Time,
			},
			want:    datetime0,
			wantErr: false,
		},
		{
			name: "正常：カスタムゼロ値",
			args: args{
				val: NilDatetime.Time(),
			},
			want:    NilDatetime,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDatetime(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDatetime() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatetime() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestInitDatetime(t *testing.T) {
	tests := []struct {
		name    string
		want    Datetime
		wantErr bool
	}{
		{
			name:    "正常",
			want:    datetime0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitDatetime()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitDatetime() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDatetime() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestDatetime_Time(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want time.Time
	}{
		{
			name: "正常",
			d:    datetime0,
			want: datetime0Time,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.Time() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestDatetime_PTime(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want *time.Time
	}{
		{
			name: "正常：not nil、値を返すこと",
			d:    datetime0,
			want: &datetime0Time,
		},
		{
			name: "正常：nil、nilを返すこと",
			d:    NilDatetime,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.PTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.PTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_TimeStringISO8601(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want string
	}{
		{
			name: "正常：utcの場合",
			d:    datetime0,
			want: datetime0StringISO8601,
		},
		{
			name: "正常：jstの場合",
			d:    jstDatetime0,
			want: jstDatetime0StringISO8601,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.TimeStringISO8601(); got != tt.want {
				t.Errorf("Datetime.TimeStringISO8601() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_PTimeStringISO8601(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want *string
	}{
		{
			name: "正常",
			d:    datetime0,
			want: &datetime0StringISO8601,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.PTimeStringISO8601(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.PTimeStringISO8601() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_Add(t *testing.T) {
	type args struct {
		val time.Duration
	}
	tests := []struct {
		name string
		d    Datetime
		args args
		want Datetime
	}{
		{
			name: "正常",
			d:    Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			args: args{
				time.Duration(time.Hour),
			},
			want: Datetime{time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Add(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_After(t *testing.T) {
	type args struct {
		val Datetime
	}
	tests := []struct {
		name string
		d    Datetime
		args args
		want bool
	}{
		{
			name: "正常：true",
			d:    Datetime{time.Date(2021, 1, 1, 0, 0, 0, 1, time.UTC)},
			args: args{
				Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			want: true,
		},
		{
			name: "正常：false、等しい場合",
			d:    Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			args: args{
				Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			want: false,
		},
		{
			name: "正常：false、前の場合",
			d:    Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			args: args{
				Datetime{time.Date(2021, 1, 1, 0, 0, 0, 1, time.UTC)},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.After(tt.args.val); got != tt.want {
				t.Errorf("Datetime.After() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_Truncate(t *testing.T) {
	type args struct {
		val time.Duration
	}
	tests := []struct {
		name string
		d    Datetime
		args args
		want Datetime
	}{
		{
			name: "正常：utc, hourの場合",
			d:    Datetime{time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)},
			args: args{
				time.Duration(time.Hour),
			},
			want: Datetime{time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)},
		},
		{
			name: "正常：jst, dayの場合",
			d:    Datetime{time.Date(2021, 1, 2, 1, 1, 1, 1, util.JST)},
			args: args{
				time.Duration(24 * time.Hour),
			},
			want: Datetime{time.Date(2021, 1, 2, 0, 0, 0, 0, util.JST)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Truncate(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_Date(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want Datetime
	}{
		{
			name: "正常：utcの場合",
			d:    Datetime{time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)},
			want: Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "正常：jstの場合",
			d:    Datetime{time.Date(2021, 1, 1, 1, 1, 1, 1, util.JST)},
			want: Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, util.JST)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Date(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_UTC(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want Datetime
	}{
		{
			name: "正常：utcの場合",
			d:    Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			want: Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "正常：jstの場合",
			d:    Datetime{time.Date(2021, 1, 1, 9, 0, 0, 0, util.JST)},
			want: Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.UTC(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.UTC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_JST(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want Datetime
	}{
		{
			name: "正常：utcの場合",
			d:    Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			want: Datetime{time.Date(2021, 1, 1, 9, 0, 0, 0, util.JST)},
		},
		{
			name: "正常：jstの場合",
			d:    Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, util.JST)},
			want: Datetime{time.Date(2021, 1, 1, 0, 0, 0, 0, util.JST)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.JST(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.JST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatetime_JSTDateInUTCDatetime(t *testing.T) {
	tests := []struct {
		name string
		d    Datetime
		want Datetime
	}{
		{
			name: "正常：utcの場合",
			d:    Datetime{time.Date(2021, 1, 1, 23, 0, 0, 0, time.UTC)},
			want: Datetime{time.Date(2021, 1, 1, 15, 0, 0, 0, time.UTC)},
		},
		{
			name: "正常：jstの場合",
			d:    Datetime{time.Date(2021, 1, 2, 0, 0, 0, 0, util.JST)},
			want: Datetime{time.Date(2021, 1, 1, 15, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.JSTDateInUTCDatetime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Datetime.JSTDateInUTCDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}
