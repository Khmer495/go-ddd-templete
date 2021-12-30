package entity

import (
	"reflect"
	"testing"

	"github.com/oklog/ulid/v2"
)

var id0String string = "01DQ4H6WA0ZPX4V3GRY7TJ0J70"
var id0ulid, _ = ulid.Parse(id0String)
var id0 = Id{id0ulid}
var id1String string = "01D0KDBRASGD5HRSNDCKA0AH53"
var id1ulid, _ = ulid.Parse(id1String)
var id1 = Id{id1ulid}
var ids0_1 = Ids{id0, id1}

func init() {
	NewUlidFunc = func() ulid.ULID {
		return id0ulid
	}
}

func TestNewId(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    Id
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: id0String,
			},
			want:    id0,
			wantErr: false,
		},
		{
			name: "異常：カスタムゼロ値",
			args: args{
				val: NilId.Ulid().String(),
			},
			want:    NilId,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewId(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewId() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewId() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestInitId(t *testing.T) {
	tests := []struct {
		name    string
		want    Id
		wantErr bool
	}{
		{
			name:    "正常",
			want:    id0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitId()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitId() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitId() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestId_Ulid(t *testing.T) {
	tests := []struct {
		name string
		i    Id
		want ulid.ULID
	}{
		{
			name: "正常",
			i:    id0,
			want: id0ulid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Ulid(); got != tt.want {
				t.Errorf("Id.Ulid() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestIds_String(t *testing.T) {
	tests := []struct {
		name string
		is   Ids
		want []string
	}{
		{
			name: "正常",
			is:   ids0_1,
			want: []string{id0String, id1String},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.is.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ids.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
