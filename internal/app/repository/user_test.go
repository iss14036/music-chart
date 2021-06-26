package repository

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/tools/driver"
	"reflect"
	"testing"
)

func TestUser_GetUser(t *testing.T) {
	type args struct {
		filter entity.UserFilter
	}
	tests := []struct {
		name    string
		args    args
		want    entity.User
		err     error
		wantErr bool
	}{
		{
			name: "Error when user not found",
			args: args{
				filter: entity.UserFilter{
					ID: "0",
				},
			},
			err:     errors.New("record not found"),
			wantErr: true,
		},
		{
			name: "Success get user",
			args: args{
				filter: entity.UserFilter{
					ID:       "1",
					Username: "test",
				},
			},
			want: entity.User{
				ID:       1,
				Username: "test",
				Password: "test",
				FullName: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			var user entity.User
			gorm := driver.NewMockGormItf(ctrl)
			gorm.EXPECT().Table("t_user").Return(gorm).AnyTimes()
			gorm.EXPECT().Select("*").Return(gorm).AnyTimes()
			gorm.EXPECT().Where("id = ?", gomock.Any()).Return(gorm).AnyTimes()
			gorm.EXPECT().Where("username = ?", gomock.Any()).Return(gorm).AnyTimes()
			gorm.EXPECT().Find(&user).
				Do(func(p *entity.User) {
					*p = tt.want
				}).Return(gorm)
			gorm.EXPECT().Error().Return(tt.err).AnyTimes()

			d := &User{db: gorm}
			got, err := d.GetUser(&tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestUser_InsertUser(t *testing.T) {
	type args struct {
		sg entity.User
	}
	tests := []struct {
		name      string
		args      args
		errSave   error
		errCommit error
		wantErr   bool
	}{
		{
			name: "Success insert User",
			args: args{
				sg: entity.User{
					Username: "test",
					Password: "test",
					FullName: "test",
					Hobby:    "test",
					Gender:   "pria",
					Address:  "test",
				},
			},
		},
		{
			name: "Failed insert user",
			args: args{
				sg: entity.User{},
			},
			errSave:   errors.New("error saved"),
			errCommit: errors.New("error commit"),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockGorm := driver.NewMockGormItf(ctrl)
			mockGorm.EXPECT().Begin().Return(mockGorm).AnyTimes()
			mockGorm.EXPECT().Table("t_user").Return(mockGorm).AnyTimes()
			mockGorm.EXPECT().Save(gomock.Any()).Return(mockGorm).AnyTimes()
			mockGorm.EXPECT().Commit().Return(mockGorm).AnyTimes()
			firstError := mockGorm.EXPECT().Error().Return(tt.errSave)
			mockGorm.EXPECT().Error().Return(tt.errCommit).After(firstError).AnyTimes()
			mockGorm.EXPECT().Rollback().AnyTimes()

			d := &User{
				db: mockGorm,
			}
			err := d.InsertUser(&tt.args.sg)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}