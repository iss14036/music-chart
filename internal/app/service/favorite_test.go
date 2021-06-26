package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/app/repository"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"reflect"
	"testing"
)

func TestFavoriteService_SetFavorite(t *testing.T) {
	type args struct {
		favorite *entity.Favorite
	}
	tests := []struct {
		name        string
		args        args
		expectMusic entity.MusicResult
		err         error
		wantErr     bool
	}{
		{
			name: "failed set favorite, music not found",
			args: args{
				&entity.Favorite{
					UserID:  1,
					MusicID: 0,
				},
			},
			err:     errors.New("error music not found"),
			wantErr: true,
		},
		{
			name: "success set favorite, return nil error",
			args: args{
				&entity.Favorite{
					UserID:  1,
					MusicID: 1,
				},
			},
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepoMusic := repository.NewMockMusicRepositoryItf(ctrl)
			mockRepoMusic.EXPECT().GetMusic(gomock.Any()).Return(tt.expectMusic, tt.err)

			mockRepo := repository.NewMockFavoriteRepositoryItf(ctrl)
			if tt.err == nil {
				mockRepo.EXPECT().SetFavorite(gomock.Any()).Return(tt.err)
			}

			d := &Favorite{
				repo:      mockRepo,
				musicRepo: mockRepoMusic,
			}

			if err := d.SetFavorite(tt.args.favorite); (err != nil) != tt.wantErr {
				t.Errorf("SetFavorite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFavoriteService_GetFavorite(t *testing.T) {
	type args struct {
		filter *entity.FavoriteFilter
	}
	tests := []struct {
		name    string
		args    args
		want    entity.FavoriteResult
		err     error
		wantErr bool
	}{
		{
			name: "Error fetching music not found for user id zero",
			args: args{
				filter: &entity.FavoriteFilter{
					UserID: "0",
					DataTable: entity.FilterDataTable{
						Sort: entity.Sort{},
						Pagination: entity.PaginationFilter{
							DisablePagination: true,
						},
					},
				},
			},
			err:     errors.New("record not found"),
			wantErr: true,
		}, {
			name: "Success Get Favorite",
			args: args{
				filter: &entity.FavoriteFilter{
					UserID: "1",
					DataTable: entity.FilterDataTable{
						Sort: entity.Sort{
							Field:     "update_at",
							Direction: "desc",
						},
						Pagination: entity.PaginationFilter{
							Page:   1,
							Limit:  10,
							Offset: 0,
						},
					},
				},
			},
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockFavoriteRepositoryItf(ctrl)
			mockRepo.EXPECT().GetFavorite(gomock.Any()).Return(tt.want, tt.err)

			d := &Favorite{
				repo: mockRepo,
			}
			got, err := d.GetFavorite(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFavorite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFavorite() got = %v, want %v", got, tt.want)
			}
		})
	}
}
