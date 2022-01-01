package model

import (
	"fmt"

	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type Page struct {
	int
}

var NilPage = Page{-1}

func NewPage(val int) (Page, error) {
	if val < 1 {
		return NilPage, cerror.NewInvalidArgumentError("val < 1", fmt.Sprintf("ページは1以上の整数を指定してください。\nhave: %d", val))
	}
	return Page{val}, nil
}

func (p Page) Int() int {
	return p.int
}
