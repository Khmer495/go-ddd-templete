package entity

import (
	"reflect"
	"testing"
	"time"
)

var user0 = User{
	id:        id0,
	createdAt: datetime0,
	name:      userName0,
}
var user1 = User{
	id:        id1,
	createdAt: datetime1,
	name:      userName1,
}
var users0 = Users{&user0}
var users0_1 = Users{&user0, &user1}

func TestNewUser(t *testing.T) {
	type args struct {
		idString      string
		createdAtTime time.Time
		nameString    string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				idString:      id0String,
				createdAtTime: datetime0Time,
				nameString:    userName0String,
			},
			want: User{
				id:        id0,
				createdAt: datetime0,
				name:      userName0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.idString, tt.args.createdAtTime, tt.args.nameString)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestInitUser(t *testing.T) {
	type args struct {
		nameString string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				nameString: userName0String,
			},
			want:    user0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitUser(tt.args.nameString)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitUser() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitUser() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestUser_Id(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want Id
	}{
		{
			name: "正常",
			u:    user0,
			want: id0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Id(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Id() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestUser_CreatedAt(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want Datetime
	}{
		{
			name: "正常",
			u:    user0,
			want: datetime0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.CreatedAt() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestUser_Name(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want UserName
	}{
		{
			name: "正常",
			u:    user0,
			want: userName0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Name() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestUser_SetName(t *testing.T) {
	u := user0
	type args struct {
		nameString string
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		want    UserName
		wantErr bool
	}{
		{
			name: "正常",
			u:    &u,
			args: args{
				nameString: u.Name().String() + "set",
			},
			want:    UserName{u.Name().String() + "set"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.u.SetName(tt.args.nameString)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.SetName() error = %+v, wantErr %+v", err, tt.wantErr)
			}
			if (err == nil) && !reflect.DeepEqual(tt.u.Name(), tt.want) {
				t.Errorf("tt.u.Name() = \n%+v\nwant\n%+v", tt.u.Name(), tt.want)
			}
		})
	}
}
