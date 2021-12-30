package entity

import (
	"time"

	"golang.org/x/xerrors"
)

type User struct {
	id        Id
	createdAt Datetime
	name      UserName
}

var NilUser = User{
	id:        NilId,
	createdAt: NilDatetime,
	name:      NilUserName,
}

type Users []*User

var NilUsers Users = nil

func NewUser(idString string, createdAtTime time.Time, nameString string) (User, error) {
	id, err := NewId(idString)
	if err != nil {
		return NilUser, xerrors.Errorf("NewId: %w", err)
	}
	createdAt, err := NewDatetime(createdAtTime)
	if err != nil {
		return NilUser, xerrors.Errorf("NewDatetime: %w", err)
	}
	name, err := NewUserName(nameString)
	if err != nil {
		return NilUser, xerrors.Errorf("NewUserName: %w", err)
	}
	return User{
		id:        id,
		createdAt: createdAt,
		name:      name,
	}, nil
}

func InitUser(nameString string) (User, error) {
	id, err := InitId()
	if err != nil {
		return NilUser, xerrors.Errorf("InitId: %w", err)
	}
	createdAt, err := InitDatetime()
	if err != nil {
		return NilUser, xerrors.Errorf("InitDatetime: %w", err)
	}
	user, err := NewUser(id.Ulid().String(), createdAt.Time(), nameString)
	if err != nil {
		return NilUser, xerrors.Errorf("NewUser: %w", err)
	}
	return user, nil
}

func (u User) Id() Id {
	return u.id
}

func (u User) CreatedAt() Datetime {
	return u.createdAt
}

func (u User) Name() UserName {
	return u.name
}

func (u *User) SetName(val string) error {
	name, err := NewUserName(val)
	if err != nil {
		return xerrors.Errorf("NewUserName: %w", err)
	}
	u.name = name
	return nil
}
