package model

import (
	"fmt"

	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type Index struct {
	int
}

var NilIndex = Index{-1}

func NewIndex(val int) (Index, error) {
	if val < 0 {
		return NilIndex, cerror.NewInvalidArgumentError("val < 0", fmt.Sprintf("インデックスは0以上の整数を指定してください。\nhave: %d", val))
	}
	return Index{val}, nil
}

func (i Index) Int() int {
	return i.int
}

func (i Index) PInt() *int {
	if i == NilIndex {
		return nil
	}
	iInt := i.Int()
	return &iInt
}
