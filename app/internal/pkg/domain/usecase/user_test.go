package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/Khmer495/go-templete/internal/pkg/domain/entity"
	"github.com/Khmer495/go-templete/internal/pkg/domain/repository"
	"github.com/Khmer495/go-templete/internal/pkg/domain/service"
	"github.com/Khmer495/go-templete/internal/pkg/util/cerror"
	"github.com/Khmer495/go-templete/test/testdata"
	gomock "github.com/golang/mock/gomock"
)

func getMocksForUser(t *testing.T) (
	*service.MockIUserService,
	*repository.MockIUserRepository,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return service.NewMockIUserService(ctrl),
		repository.NewMockIUserRepository(ctrl)

}

func TestNewUserUsecase(t *testing.T) {
	us, ur := getMocksForUser(t)
	type args struct {
		us service.IUserService
		ur repository.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want IUserUsecase
	}{
		{
			name: "正常",
			args: args{
				us: us,
				ur: ur,
			},
			want: userUsecase{
				us: us,
				ur: ur,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.us, tt.args.ur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_Register(t *testing.T) {
	type args struct {
		name string
	}
	type mockVals struct {
		usSelfIdRessErr    error
		urRegisterArgsUser entity.User
	}
	tests := []struct {
		name     string
		args     args
		want     entity.User
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常",
			args: args{
				name: testdata.UserName0String,
			},
			want:    testdata.User0,
			wantErr: false,
			mockVals: mockVals{
				usSelfIdRessErr:    cerror.NewNotFoundError("", ""),
				urRegisterArgsUser: testdata.User0,
			},
		},
		{
			name: "異常：既にユーザーが登録されている（トークンがDBに存在する）場合、ユーザー登録を受け付けない",
			args: args{
				name: testdata.UserName0String,
			},
			want:    entity.NilUser,
			wantErr: true,
			mockVals: mockVals{
				usSelfIdRessErr: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			us, ur := getMocksForUser(t)
			us.EXPECT().SelfId(ctx).Return(entity.NilId, tt.mockVals.usSelfIdRessErr).Times(1)
			ur.EXPECT().Register(ctx, tt.mockVals.urRegisterArgsUser).Return(nil).Times(1)
			uu := NewUserUsecase(us, ur)
			got, err := uu.Register(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetSelfProfile(t *testing.T) {
	type mockVals struct {
		urSelfRessUser entity.User
	}
	tests := []struct {
		name     string
		want     entity.User
		wantErr  bool
		mockVals mockVals
	}{
		{
			name:    "正常",
			want:    testdata.User0,
			wantErr: false,
			mockVals: mockVals{
				urSelfRessUser: testdata.User0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			us, ur := getMocksForUser(t)
			ur.EXPECT().Self(ctx).Return(tt.mockVals.urSelfRessUser, nil).Times(1)
			uu := NewUserUsecase(us, ur)
			got, err := uu.GetSelfProfile(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetSelfProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetSelfProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetList(t *testing.T) {
	type args struct {
		limitInt    int
		pageInt     int
		pNameString *string
	}
	type mockVals struct {
		urListTimes                   int
		urListArgsLimit               entity.Limit
		urListArgsPage                entity.Page
		urListRessUsers               entity.Users
		urSearchByNamePrefixTimes     int
		urSearchByNamePrefixArgsLimit entity.Limit
		urSearchByNamePrefixArgsPage  entity.Page
		urSearchByNamePrefixArgsName  entity.UserName
		urSearchByNamePrefixRessUsers entity.Users
	}
	tests := []struct {
		name     string
		args     args
		want     entity.Users
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常：userNameがnilの場合、ur.Listを呼ぶこと",
			args: args{
				limitInt:    testdata.Limit0Int,
				pageInt:     testdata.Page0Int,
				pNameString: nil,
			},
			want:    testdata.Users0_1,
			wantErr: false,
			mockVals: mockVals{
				urListTimes:               1,
				urListArgsLimit:           testdata.Limit0,
				urListArgsPage:            testdata.Page0,
				urListRessUsers:           testdata.Users0_1,
				urSearchByNamePrefixTimes: 0,
			},
		},
		{
			name: "正常：userNameがnot nilの場合、ur.SearchByNamePrefixを呼ぶこと",
			args: args{
				limitInt:    testdata.Limit0Int,
				pageInt:     testdata.Page0Int,
				pNameString: &testdata.UserName0String,
			},
			want:    testdata.Users0_1,
			wantErr: false,
			mockVals: mockVals{
				urListTimes:                   0,
				urSearchByNamePrefixTimes:     1,
				urSearchByNamePrefixArgsLimit: testdata.Limit0,
				urSearchByNamePrefixArgsPage:  testdata.Page0,
				urSearchByNamePrefixArgsName:  testdata.UserName0,
				urSearchByNamePrefixRessUsers: testdata.Users0_1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			us, ur := getMocksForUser(t)
			ur.EXPECT().List(ctx, tt.mockVals.urListArgsLimit, tt.mockVals.urListArgsPage).Return(tt.mockVals.urListRessUsers, nil).Times(tt.mockVals.urListTimes)
			ur.EXPECT().SearchByNamePrefix(ctx, tt.mockVals.urSearchByNamePrefixArgsLimit, tt.mockVals.urSearchByNamePrefixArgsPage, tt.mockVals.urSearchByNamePrefixArgsName).Return(tt.mockVals.urSearchByNamePrefixRessUsers, nil).Times(tt.mockVals.urSearchByNamePrefixTimes)
			uu := NewUserUsecase(us, ur)
			got, err := uu.GetList(ctx, tt.args.limitInt, tt.args.pageInt, tt.args.pNameString)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_ChangeSelfProfile(t *testing.T) {
	type args struct {
		pName *string
	}
	type mockVals struct {
		urSelfRessUser   entity.User
		urChangeArgsUser entity.User
	}
	tests := []struct {
		name     string
		args     args
		want     entity.User
		wantErr  bool
		mockVals mockVals
	}{
		{
			name: "正常",
			args: args{
				pName: &testdata.UserName1String,
			},
			want:    testdata.User0v2,
			wantErr: false,
			mockVals: mockVals{
				urSelfRessUser:   testdata.User0,
				urChangeArgsUser: testdata.User0v2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			us, ur := getMocksForUser(t)
			ur.EXPECT().Self(ctx).Return(tt.mockVals.urSelfRessUser, nil).Times(1)
			ur.EXPECT().Change(ctx, tt.mockVals.urChangeArgsUser).Return(nil).Times(1)
			uu := NewUserUsecase(us, ur)
			got, err := uu.ChangeSelfProfile(ctx, tt.args.pName)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.ChangeSelfProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.ChangeSelfProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
