package entity

import (
	"reflect"
	"strings"
	"testing"
)

var teamDescription0String = "aA0ａA０あアｱ阿!# ！＃　\n"
var teamDescription0 = TeamDescription{teamDescription0String}

func TestNewTeamDescription(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    TeamDescription
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: teamDescription0String,
			},
			want:    teamDescription0,
			wantErr: false,
		},
		{
			name: "異常：カスタムゼロ値",
			args: args{
				val: NilTeamDescription.String(),
			},
			want:    NilTeamDescription,
			wantErr: true,
		},
		{
			name: "検証：256文字以下許可",
			args: args{
				val: strings.Repeat("1234567890", 25) + "123456",
			},
			want:    TeamDescription{strings.Repeat("1234567890", 25) + "123456"},
			wantErr: false,
		},
		{
			name: "検証：257文字以上禁止",
			args: args{
				val: strings.Repeat("1234567890", 25) + "1234567",
			},
			want:    NilTeamDescription,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTeamDescription(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTeamDescription() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTeamDescription() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestTeamDescription_String(t *testing.T) {
	tests := []struct {
		name string
		td   TeamDescription
		want string
	}{
		{
			name: "正常",
			td:   teamDescription0,
			want: teamDescription0String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.td.String(); got != tt.want {
				t.Errorf("TeamDescription.String() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}
