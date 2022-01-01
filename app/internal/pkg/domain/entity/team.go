package model

import (
	"time"

	"golang.org/x/xerrors"
)

type Team struct {
	id          Id
	createdAt   Datetime
	createUser  User
	name        TeamName
	description TeamDescription
	users       Users
}

var NilTeam = Team{
	id:          NilId,
	createdAt:   NilDatetime,
	createUser:  NilUser,
	name:        NilTeamName,
	description: NilTeamDescription,
	users:       NilUsers,
}

type Teams []*Team

var NilTeams Teams = nil

func NewTeam(idString string, createdAtTime time.Time, createUser User, nameString string, descriptionString string, users Users) (Team, error) {
	id, err := NewId(idString)
	if err != nil {
		return NilTeam, xerrors.Errorf("NewId: %w", err)
	}
	createdAt, err := NewDatetime(createdAtTime)
	if err != nil {
		return NilTeam, xerrors.Errorf("NewDatetim: %w", err)
	}
	name, err := NewTeamName(nameString)
	if err != nil {
		return NilTeam, xerrors.Errorf("NewTeamName: %w", err)
	}
	description, err := NewTeamDescription(descriptionString)
	if err != nil {
		return NilTeam, xerrors.Errorf("NewTeamDescription: %w", err)
	}
	return Team{
		id:          id,
		createdAt:   createdAt,
		createUser:  createUser,
		name:        name,
		description: description,
		users:       users,
	}, nil
}

func InitTeam(createUser User, nameString string, descriptionString string) (Team, error) {
	id, err := InitId()
	if err != nil {
		return NilTeam, xerrors.Errorf("InitId: %w", err)
	}
	createdAt, err := InitDatetime()
	if err != nil {
		return NilTeam, xerrors.Errorf("InitDatetime: %w", err)
	}
	team, err := NewTeam(id.Ulid().String(), createdAt.Time(), createUser, nameString, descriptionString, Users{&createUser})
	if err != nil {
		return NilTeam, xerrors.Errorf("NewTeam: %w", err)
	}
	return team, nil
}

func (t Team) Id() Id {
	return t.id
}

func (t Team) CreatedAt() Datetime {
	return t.createdAt
}

func (t Team) CreateUser() User {
	return t.createUser
}

func (t Team) Name() TeamName {
	return t.name
}

func (t *Team) SetName(val string) error {
	name, err := NewTeamName(val)
	if err != nil {
		return xerrors.Errorf("NewTeamName: %w", err)
	}
	t.name = name
	return nil
}

func (t Team) Description() TeamDescription {
	return t.description
}

func (t *Team) SetDesctiprion(val string) error {
	description, err := NewTeamDescription(val)
	if err != nil {
		return xerrors.Errorf("NewTeamDescription: %w", err)
	}
	t.description = description
	return nil
}

func (t Team) Users() Users {
	return t.users
}

func (t *Team) SetUsers(val Users) error {
	t.users = val
	return nil
}

func (t Team) IsUser(userId Id) bool {
	for _, user := range t.users {
		if user.id == userId {
			return true
		}
	}
	return false
}
