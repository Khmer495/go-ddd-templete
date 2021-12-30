package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/test/testdata"
	"github.com/golang/mock/gomock"
)

func getMocksForTeam(t *testing.T) (
	*service.MockITeamService,
	*service.MockIUserService,
	*repository.MockITeamRepository,
	*repository.MockIUserRepository,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return service.NewMockITeamService(ctrl),
		service.NewMockIUserService(ctrl),
		repository.NewMockITeamRepository(ctrl),
		repository.NewMockIUserRepository(ctrl)
}

func TestNewTeamUsecase(t *testing.T) {
	ts, us, tr, ur := getMocksForTeam(t)
	type args struct {
		ts service.ITeamService
		us service.IUserService
		tr repository.ITeamRepository
		ur repository.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want ITeamUsecase
	}{
		{
			name: "正常",
			args: args{
				ts: ts,
				us: us,
				tr: tr,
				ur: ur,
			},
			want: teamUsecase{
				ts: ts,
				us: us,
				tr: tr,
				ur: ur,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTeamUsecase(tt.args.ts, tt.args.us, tt.args.tr, tt.args.ur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTeamUsecase() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_Create(t *testing.T) {
	type args struct {
		name        string
		description string
	}
	type mockVals struct {
		urSelfRessUser   entity.User
		trCreateArgsTeam entity.Team
	}
	tests := []struct {
		name     string
		args     args
		want     entity.Team
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常",
			args: args{
				name:        testdata.TeamName0String,
				description: testdata.TeamDescription0String,
			},
			want:    testdata.Team0v3,
			wantErr: false,
			mockVals: mockVals{
				urSelfRessUser:   testdata.User0,
				trCreateArgsTeam: testdata.Team0v3,
			},
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		ts, us, tr, ur := getMocksForTeam(t)
		ur.EXPECT().Self(ctx).Return(tt.mockVals.urSelfRessUser, nil).Times(1)
		tr.EXPECT().Create(ctx, tt.mockVals.trCreateArgsTeam).Return(nil).Times(1)
		tu := NewTeamUsecase(ts, us, tr, ur)
		t.Run(tt.name, func(t *testing.T) {
			got, err := tu.Create(ctx, tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.Create() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsecase.Create() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_GetList(t *testing.T) {
	type args struct {
		limitInt    int
		pageInt     int
		pNameString *string
	}
	type mockVals struct {
		trListTimes                   int
		trListArgsLimit               entity.Limit
		trListArgsPage                entity.Page
		trListRessTeams               entity.Teams
		trSearchByNamePrefixTimes     int
		trSearchByNamePrefixArgsLimit entity.Limit
		trSearchByNamePrefixArgsPage  entity.Page
		trSearchByNamePrefixArgsName  entity.TeamName
		trSearchByNamePrefixRessTeams entity.Teams
	}
	tests := []struct {
		name     string
		args     args
		want     entity.Teams
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常：teamNameがnilの場合、tr.Listを呼ぶこと",
			args: args{
				limitInt:    testdata.Limit0Int,
				pageInt:     testdata.Page0Int,
				pNameString: nil,
			},
			want:    testdata.Teams0,
			wantErr: false,
			mockVals: mockVals{
				trListTimes:               1,
				trListArgsLimit:           testdata.Limit0,
				trListArgsPage:            testdata.Page0,
				trListRessTeams:           testdata.Teams0,
				trSearchByNamePrefixTimes: 0,
			},
		},
		{
			name: "正常：teamNameがnot nilの場合、tr.SearchByNamePrefixを呼ぶこと",
			args: args{
				limitInt:    testdata.Limit0Int,
				pageInt:     testdata.Page0Int,
				pNameString: &testdata.TeamName0String,
			},
			want:    testdata.Teams0,
			wantErr: false,
			mockVals: mockVals{
				trListTimes:                   0,
				trSearchByNamePrefixTimes:     1,
				trSearchByNamePrefixArgsLimit: testdata.Limit0,
				trSearchByNamePrefixArgsPage:  testdata.Page0,
				trSearchByNamePrefixArgsName:  testdata.TeamName0,
				trSearchByNamePrefixRessTeams: testdata.Teams0,
			},
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		ts, us, tr, ur := getMocksForTeam(t)
		tr.EXPECT().List(ctx, tt.mockVals.trListArgsLimit, tt.mockVals.trListArgsPage).Return(tt.mockVals.trListRessTeams, nil).Times(tt.mockVals.trListTimes)
		tr.EXPECT().SearchByNamePrefix(ctx, tt.mockVals.trSearchByNamePrefixArgsLimit, tt.mockVals.trSearchByNamePrefixArgsPage, tt.mockVals.trSearchByNamePrefixArgsName).Return(tt.mockVals.trSearchByNamePrefixRessTeams, nil).Times(tt.mockVals.trSearchByNamePrefixTimes)
		tu := NewTeamUsecase(ts, us, tr, ur)
		t.Run(tt.name, func(t *testing.T) {
			got, err := tu.GetList(ctx, tt.args.limitInt, tt.args.pageInt, tt.args.pNameString)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.GetList() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsecase.GetList() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_Change(t *testing.T) {
	type args struct {
		teamIdString string
		pName        *string
		pDescription *string
	}
	type mockVals struct {
		trOneArgsTeamId     entity.Id
		trOneRessTeam       entity.Team
		urSelectArgsUserIds entity.Ids
		urSelectRessUsers   entity.Users
		urSelfRessUser      entity.User
		trChangeArgsTeam    entity.Team
	}
	tests := []struct {
		name     string
		args     args
		want     entity.Team
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常",
			args: args{
				teamIdString: testdata.Id0String,
				pName:        &testdata.TeamName1String,
				pDescription: &testdata.TeamDescription1String,
			},
			want:    testdata.Team0v2,
			wantErr: false,
			mockVals: mockVals{
				trOneArgsTeamId:     testdata.Id0,
				trOneRessTeam:       testdata.Team0,
				urSelectArgsUserIds: testdata.Ids0_1,
				urSelectRessUsers:   testdata.Users0_1,
				urSelfRessUser:      testdata.User0,
				trChangeArgsTeam:    testdata.Team0v2,
			},
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		ts, us, tr, ur := getMocksForTeam(t)
		tr.EXPECT().One(ctx, tt.mockVals.trOneArgsTeamId).Return(tt.mockVals.trOneRessTeam, nil).Times(1)
		ur.EXPECT().Select(ctx, entity.NilLimit, entity.NilPage, tt.mockVals.urSelectArgsUserIds).Return(tt.mockVals.urSelectRessUsers, nil).Times(1)
		ur.EXPECT().Self(ctx).Return(tt.mockVals.urSelfRessUser, nil).Times(1)
		tr.EXPECT().Change(ctx, tt.mockVals.trChangeArgsTeam).Return(nil).Times(1)
		tu := NewTeamUsecase(ts, us, tr, ur)
		t.Run(tt.name, func(t *testing.T) {
			got, err := tu.Change(ctx, tt.args.teamIdString, tt.args.pName, tt.args.pDescription)
			if (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.Change() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("teamUsecase.Change() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_teamUsecase_Delete(t *testing.T) {
	type args struct {
		teamIdString string
	}
	type mockVals struct {
		trOneArgsTeamId    entity.Id
		trOneRessTeam      entity.Team
		trDeleteArgsTeamId entity.Id
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常",
			args: args{
				teamIdString: testdata.Id0String,
			},
			wantErr: false,
			mockVals: mockVals{
				trOneArgsTeamId:    testdata.Id0,
				trOneRessTeam:      testdata.Team0,
				trDeleteArgsTeamId: testdata.Id0,
			},
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		ts, us, tr, ur := getMocksForTeam(t)
		tr.EXPECT().One(ctx, tt.mockVals.trOneArgsTeamId).Return(tt.mockVals.trOneRessTeam, nil).Times(1)
		tr.EXPECT().Delete(ctx, tt.mockVals.trDeleteArgsTeamId).Return(nil).Times(1)
		tu := NewTeamUsecase(ts, us, tr, ur)
		t.Run(tt.name, func(t *testing.T) {
			if err := tu.Delete(ctx, tt.args.teamIdString); (err != nil) != tt.wantErr {
				t.Errorf("teamUsecase.Delete() error = %+v, wantErr %+v", err, tt.wantErr)
			}
		})
	}
}

func Test_teamUsecase_Join(t *testing.T) {
	type args struct {
		teamIdString string
	}
	type mockVals struct {
		trOneArgsTeamId    entity.Id
		trOneRessTeam      entity.Team
		usSelfIdRessUserId entity.Id
		trJoinArgsTeamId   entity.Id
		trJoinArgsUserId   entity.Id
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常",
			args: args{
				teamIdString: testdata.Id0String,
			},
			wantErr: false,
			mockVals: mockVals{
				trOneArgsTeamId:    testdata.Id0,
				trOneRessTeam:      testdata.Team0,
				usSelfIdRessUserId: testdata.Id2,
				trJoinArgsTeamId:   testdata.Id0,
				trJoinArgsUserId:   testdata.Id2,
			},
		},
		{
			name: "異常：加入済みであるときエラーを返す",
			args: args{
				teamIdString: testdata.Id0String,
			},
			wantErr: true,
			mockVals: mockVals{
				trOneArgsTeamId:    testdata.Id0,
				trOneRessTeam:      testdata.Team0,
				usSelfIdRessUserId: testdata.Id0,
			},
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		ts, us, tr, ur := getMocksForTeam(t)
		tr.EXPECT().One(ctx, tt.mockVals.trOneArgsTeamId).Return(tt.mockVals.trOneRessTeam, nil).Times(1)
		us.EXPECT().SelfId(ctx).Return(tt.mockVals.usSelfIdRessUserId, nil).Times(1)
		tr.EXPECT().Join(ctx, tt.mockVals.trJoinArgsTeamId, tt.mockVals.trJoinArgsUserId).Return(nil).Times(1)
		tu := NewTeamUsecase(ts, us, tr, ur)
		t.Run(tt.name, func(t *testing.T) {
			if err := tu.Join(ctx, tt.args.teamIdString); (err != nil) != tt.wantErr {
				t.Errorf("tu.Join() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
		})
	}
}
