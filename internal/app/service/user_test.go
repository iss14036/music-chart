package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/app/repository"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/internal/pkg/gateway"
	"reflect"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {
	type args struct {
		filter entity.UserFilter
	}
	tests := []struct {
		name     string
		args     args
		want    entity.User
		err error
		wantErr bool
	} {
		{
			name: "Error fetching user not found",
			args: args{
				filter: entity.UserFilter{
					ID:       "0",
				},
			},
			err:    errors.New("record not found"),
			wantErr: true,
		}, {
			name: "Success Get Ticker",
			args: args{
				filter: entity.UserFilter{
					ID:       "1",
				},
			},
			want: entity.User{
				ID:       1,
				Username: "test",
				Password: "test",
				FullName: "test",
				Hobby:    "test",
				Gender:   "pria",
				Address:  "test",
			},
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockUserRepositoryItf(ctrl)
			mockRepo.EXPECT().GetUser(gomock.Any()).Return(tt.want, tt.err)

			d := &User{
				repo: mockRepo,
			}
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

func TestUserService_InsertUser(t *testing.T) {
	type args struct {
		user entity.User
	}
	tests := []struct {
		name     string
		args     args
		expected error
		wantErr  bool
	}{
		{
			name: "Success insert user, return nil",
			args: args{
				user: entity.User{
					Username: "test",
					Password: "test",
					FullName: "test",
					Hobby:    "test",
					Gender:   "test",
					Address:  "test",
				},
			},
		},
		{
			name: "Error insert user, return error",
			args: args{
				user: entity.User{},
			},
			expected: errors.New("error insert user"),
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockUserRepositoryItf(ctrl)
			mockRepo.EXPECT().InsertUser(gomock.Any()).Return(tt.expected)

			d := &User{
				repo: mockRepo,
			}
			if err := d.InsertUser(&tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type args struct {
		filter entity.UserFilter
	}
	type expected struct {
		token string
		err error
	}
	type expectedGetUser struct {
		user entity.User
		err error
	}
	tests := []struct {
		name     string
		args     args
		expectedGetUser  expectedGetUser
		expected expected
		wantErr  bool
	}{
		{
			name: "Success login, return token",
			args: args{
				filter: entity.UserFilter{
					Username: "test",
					Password: "test",
				},
			},
			expectedGetUser: expectedGetUser{
				user: entity.User{
					ID:       1,
					Username: "test",
					Password: "$2a$10$YCxWOrJyques/nPGRPILX.p8kz8WsbFi8.Xzl7S8ly9r9K..YQXHi",
				},
				err:  nil,
			},
			expected: expected{
				token: "test",
				err:   nil,
			},
		},
		{
			name: "Error login, return error",
			args: args{
				filter: entity.UserFilter{
					Username: "",
					Password: "",
				},
			},
			expectedGetUser: expectedGetUser{
				user: entity.User{},
				err:  errors.New("error login"),
			},
			expected: expected{
				token: "",
				err:   errors.New("error login"),
			},
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockUserRepositoryItf(ctrl)
			mockRepo.EXPECT().GetUser(gomock.Any()).Return(tt.expectedGetUser.user, tt.expectedGetUser.err)
			mockAuth := gateway.NewMockAuthItf(ctrl)
			if tt.expectedGetUser.err == nil {
				mockAuth.EXPECT().GetToken(gomock.Any()).Return(tt.expected.token, tt.expected.err)
			}

			d := &User{
				repo: mockRepo,
				auth: mockAuth,
			}
			got, err := d.Login(&tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected.token) {
				t.Errorf("Login() got = %v, want %v", got, tt.expected.token)
			}
		})
	}
}