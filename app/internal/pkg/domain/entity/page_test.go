package entity

import (
	"reflect"
	"testing"
)

var page1Int = 10
var page1 = Page{page1Int}

func TestNewPage(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		args    args
		want    Page
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				val: page1Int,
			},
			want:    page1,
			wantErr: false,
		},
		{
			name: "異常：カスタムゼロ値",
			args: args{
				val: NilPage.Int(),
			},
			want:    NilPage,
			wantErr: true,
		},
		{
			name: "検証：0以下禁止",
			args: args{
				val: 0,
			},
			want:    NilPage,
			wantErr: true,
		},
		{
			name: "検証：1以上許可",
			args: args{
				val: 1,
			},
			want:    Page{1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPage(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_Int(t *testing.T) {
	tests := []struct {
		name string
		p    Page
		want int
	}{
		{
			name: "正常",
			p:    page1,
			want: page1Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Int(); got != tt.want {
				t.Errorf("Page.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
