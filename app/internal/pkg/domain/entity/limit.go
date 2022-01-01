package model

import (
	"fmt"

	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
)

type Limit struct {
	int
}

var NilLimit = Limit{-1}

func NewLimit(val int) (Limit, error) {
	if val < 1 {
		return NilLimit, cerror.NewInvalidArgumentError("val < 1", fmt.Sprintf("1ページあたりの表示数は1以上の整数を指定してください。\nhave: %d", val))
	}
	return Limit{val}, nil
}

func (l Limit) Int() int {
	return l.int
}
