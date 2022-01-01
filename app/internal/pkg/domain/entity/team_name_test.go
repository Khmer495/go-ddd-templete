package model

import (
	"reflect"
	"testing"
)

var teamName0String = "aA0ａA０あアｱ阿! ！　"
var teamName0 = TeamName{teamName0String}

func TestNewTeamName(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    TeamName
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: teamName0String,
			},
			want:    teamName0,
			wantErr: false,
		},
		{
			name: "異常：カスタムゼロ値",
			args: args{
				val: NilTeamName.String(),
			},
			want:    NilTeamName,
			wantErr: true,
		},
		{
			name: "検証：空文字禁止",
			args: args{
				val: "",
			},
			want:    NilTeamName,
			wantErr: true,
		},
		{
			name: "検証：1文字以上許可",
			args: args{
				val: "1",
			},
			want:    TeamName{"1"},
			wantErr: false,
		},
		{
			name: "検証：16文字以下許可",
			args: args{
				val: "1234567890123456",
			},
			want:    TeamName{"1234567890123456"},
			wantErr: false,
		},
		{
			name: "検証：17文字以上禁止",
			args: args{
				val: "12345678901234567",
			},
			want:    NilTeamName,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTeamName(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTeamName() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTeamName() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestTeamName_String(t *testing.T) {
	tests := []struct {
		name string
		tn   TeamName
		want string
	}{
		{
			name: "正常",
			tn:   teamName0,
			want: teamName0String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tn.String(); got != tt.want {
				t.Errorf("TeamName.String() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}
