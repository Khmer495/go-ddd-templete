package entity

import (
	"reflect"
	"testing"
)

var index0Int = 10
var index0 = Index{index0Int}

func TestNewIndex(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		args    args
		want    Index
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: index0Int,
			},
			want:    index0,
			wantErr: false,
		},
		{
			name: "異常：カスタムゼロ値",
			args: args{
				val: NilIndex.Int(),
			},
			want:    NilIndex,
			wantErr: true,
		},
		{
			name: "検証：-1以下禁止",
			args: args{
				val: -1,
			},
			want:    NilIndex,
			wantErr: true,
		},
		{
			name: "検証：0以上許可",
			args: args{
				val: 0,
			},
			want:    Index{0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIndex(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Int(t *testing.T) {
	tests := []struct {
		name string
		i    Index
		want int
	}{
		{
			name: "正常",
			i:    index0,
			want: index0Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Int(); got != tt.want {
				t.Errorf("Index.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_PInt(t *testing.T) {
	tests := []struct {
		name string
		i    Index
		want *int
	}{
		{
			name: "正常：not nil、値を返すこと",
			i:    index0,
			want: &index0Int,
		},
		{
			name: "正常：nil、nilを返すこと",
			i:    NilIndex,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.PInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Index.PInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
