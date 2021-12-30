package entity

import (
	"reflect"
	"testing"
	"time"
)

var team0 = Team{
	id:          id0,
	createdAt:   datetime0,
	createUser:  user0,
	name:        teamName0,
	description: teamDescription0,
	users:       users0_1,
}

func TestNewTeam(t *testing.T) {
	type args struct {
		idString          string
		createdAtTime     time.Time
		createUser        User
		nameString        string
		descriptionString string
		users             Users
	}
	tests := []struct {
		name    string
		args    args
		want    Team
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				idString:          id0String,
				createdAtTime:     datetime0Time,
				createUser:        user0,
				nameString:        teamName0String,
				descriptionString: teamDescription0String,
				users:             users0_1,
			},
			want:    team0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTeam(tt.args.idString, tt.args.createdAtTime, tt.args.createUser, tt.args.nameString, tt.args.descriptionString, tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTeam() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTeam() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestInitTeam(t *testing.T) {
	type args struct {
		createUser        User
		nameString        string
		descriptionString string
		users             Users
	}
	tests := []struct {
		name    string
		args    args
		want    Team
		wantErr bool
	}{
		{
			name: "正常",
			args: args{
				createUser:        user0,
				nameString:        teamName0String,
				descriptionString: teamDescription0String,
				users:             users0,
			},
			want: Team{
				id:          id0,
				createdAt:   datetime0,
				createUser:  user0,
				name:        teamName0,
				description: teamDescription0,
				users:       users0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitTeam(tt.args.createUser, tt.args.nameString, tt.args.descriptionString)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitTeam() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitTeam() = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}

func TestTeam_Id(t *testing.T) {
	tests := []struct {
		name string
		tr   Team
		want Id
	}{
		{
			name: "正常",
			tr:   team0,
			want: id0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Id(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Team.Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_CreatedAt(t *testing.T) {
	tests := []struct {
		name string
		tr   Team
		want Datetime
	}{
		{
			name: "正常",
			tr:   team0,
			want: datetime0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Team.CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_CreateUser(t *testing.T) {
	tests := []struct {
		name string
		tr   Team
		want User
	}{
		{
			name: "正常",
			tr:   team0,
			want: user0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.CreateUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Team.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_Name(t *testing.T) {
	tests := []struct {
		name string
		tr   Team
		want TeamName
	}{
		{
			name: "正常",
			tr:   team0,
			want: teamName0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Team.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_SetName(t *testing.T) {
	tr := team0
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		tr      *Team
		args    args
		want    TeamName
		wantErr bool
	}{
		{
			name: "正常",
			tr:   &tr,
			args: args{
				val: tr.Name().String() + "se",
			},
			want:    TeamName{tr.Name().String() + "se"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tr.SetName(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Team.SetName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err == nil) && !reflect.DeepEqual(tt.tr.Name(), tt.want) {
				t.Errorf("tt.tr.Name() = \n%+v\nwant\n%+v", tt.tr.Name(), tt.want)
			}
		})
	}
}

func TestTeam_Description(t *testing.T) {
	tests := []struct {
		name string
		tr   Team
		want TeamDescription
	}{
		{
			name: "正常",
			tr:   team0,
			want: teamDescription0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Description(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Team.Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_SetDesctiprion(t *testing.T) {
	tr := team0
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		tr      *Team
		args    args
		want    TeamDescription
		wantErr bool
	}{
		{
			name: "正常",
			tr:   &tr,
			args: args{
				val: tr.Description().String() + "set",
			},
			want:    TeamDescription{tr.Description().String() + "set"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tr.SetDesctiprion(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Team.SetDesctiprion() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err == nil) && !reflect.DeepEqual(tt.tr.Description(), tt.want) {
				t.Errorf("tt.tr.Description() = \n%+v\nwant\n%+v", tt.tr.Description(), tt.want)
			}
		})
	}
}

func TestTeam_Users(t *testing.T) {
	tests := []struct {
		name string
		tr   Team
		want Users
	}{
		{
			name: "正常",
			tr:   team0,
			want: users0_1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Users(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Team.Users() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeam_SetUsers(t *testing.T) {
	tr := team0
	type args struct {
		val Users
	}
	tests := []struct {
		name    string
		tr      *Team
		args    args
		want    Users
		wantErr bool
	}{
		{
			name: "正常",
			tr:   &tr,
			args: args{
				val: users0,
			},
			want:    Users{&user0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tr.SetUsers(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Team.SetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err == nil) && !reflect.DeepEqual(tt.tr.Users(), tt.want) {
				t.Errorf("tt.tr.Users() = \n%+v\nwant\n%+v", tt.tr.Users(), tt.want)
			}
		})
	}
}
