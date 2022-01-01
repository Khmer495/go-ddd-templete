package model

import (
	"testing"
)

var userName0String = "aA0ａA０あアｱ阿"
var userName0 = UserName{userName0String}
var userName1String = "bB1ｂＢ１いイｲ伊"
var userName1 = UserName{userName1String}

func TestNewUserName(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    UserName
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: userName0String,
			},
			want:    userName0,
			wantErr: false,
		},
		{
			name: "異常：カスタムゼロ値",
			args: args{
				val: NilUserName.String(),
			},
			want:    NilUserName,
			wantErr: true,
		},
		{
			name: "検証：空文字禁止",
			args: args{
				val: "",
			},
			want:    NilUserName,
			wantErr: true,
		},
		{
			name: "検証：1文字以上OK",
			args: args{
				val: "1",
			},
			want:    UserName{"1"},
			wantErr: false,
		},
		{
			name: "検証：16文字以下OK",
			args: args{
				val: "1234567890123456",
			},
			want:    UserName{"1234567890123456"},
			wantErr: false,
		},
		{
			name: "検証：17文字以上禁止",
			args: args{
				val: "12345678901234567",
			},
			want:    NilUserName,
			wantErr: true,
		},
		{
			name: "検証：前後のスペースは削除",
			args: args{
				val: " 123456　",
			},
			want:    UserName{"123456"},
			wantErr: false,
		},
		{
			name: "検証：半角スペースは全角スペースに置換",
			args: args{
				val: "first　family",
			},
			want:    UserName{"first family"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserName(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserName() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewUserName() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestUserName_String(t *testing.T) {
	tests := []struct {
		name string
		un   UserName
		want string
	}{
		{
			name: "正常",
			un:   userName0,
			want: userName0String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.un.String(); got != tt.want {
				t.Errorf("UserName.String() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}
