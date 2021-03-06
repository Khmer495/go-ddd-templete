package testdata

import (
	"github.com/Khmer495/go-templete/internal/pkg/domain/model"
	"github.com/oklog/ulid/v2"
)

func init() {
	model.NewUlidFunc = func() ulid.ULID {
		return Id0.Ulid()
	}
}

var Id0String = "01DQ4H6WA0ZPX4V3GRY7TJ0J70"
var Id0, _ = model.NewId(Id0String)
var Id1String = "01D0KDBRASGD5HRSNDCKA0AH53"
var Id1, _ = model.NewId(Id1String)
var Id2String = "01F8N6AFD2Z85C0N1F4YAX8PYG"
var Id2, _ = model.NewId(Id2String)
var Ids0 = model.Ids{Id0}
var Ids0_1 = model.Ids{Id0, Id1}
