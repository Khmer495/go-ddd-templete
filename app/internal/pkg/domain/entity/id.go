package entity

import (
	"math/rand"
	"time"

	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/oklog/ulid/v2"
	"golang.org/x/xerrors"
)

type Id struct {
	ulid ulid.ULID
}

var NilId = Id{ulid.ULID{}}

type Ids []Id

var NilIds Ids = nil

func NewId(val string) (Id, error) {
	idUlid, err := ulid.ParseStrict(val)
	if err != nil {
		return NilId, cerror.NewInvalidArgumentError("idUlid.ParseStrict", "Idの形式が間違っています。")
	}
	// NilId.Ulid()(=ulid.ULID{})はulid.ParseStrictを満たすため、別途エラー処理をする
	if idUlid == NilId.Ulid() {
		return NilId, cerror.NewInvalidArgumentError("idUlid == NilId.Ulid()", "Idの形式が間違っています。")
	}
	return Id{idUlid}, nil
}

var NewUlidFunc = func() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

func InitId() (Id, error) {
	ulid := NewUlidFunc()
	id, err := NewId(ulid.String())
	if err != nil {
		return NilId, xerrors.Errorf("NewId: %w", err)
	}
	return id, nil
}

func (i Id) Ulid() ulid.ULID {
	return i.ulid
}

func (is Ids) String() []string {
	idsString := []string{}
	for _, i := range is {
		idsString = append(idsString, i.Ulid().String())
	}
	return idsString
}
