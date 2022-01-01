package model

import (
	"reflect"
	"testing"
)

var limit0Int = 10
var limit0 = Limit{limit0Int}

func TestNewLimit(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		args    args
		want    Limit
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: limit0Int,
			},
			want:    limit0,
			wantErr: false,
		},
		{
			name: "異常：カスタムゼロ値",
			args: args{
				val: NilLimit.Int(),
			},
			want:    NilLimit,
			wantErr: true,
		},
		{
			name: "検証：0以下禁止",
			args: args{
				val: 0,
			},
			want:    NilLimit,
			wantErr: true,
		},
		{
			name: "検証：1以上許可",
			args: args{
				val: 1,
			},
			want:    Limit{1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLimit(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLimit_Int(t *testing.T) {
	tests := []struct {
		name string
		l    Limit
		want int
	}{
		{
			name: "正常",
			l:    limit0,
			want: limit0Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Int(); got != tt.want {
				t.Errorf("Limit.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
