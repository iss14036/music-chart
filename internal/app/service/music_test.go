package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/app/repository"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"reflect"
	"testing"
)

func TestMusicService_FetchMusic(t *testing.T) {
	type args struct {
		filter entity.MusicFilter
	}
	tests := []struct {
		name    string
		args    args
		want    entity.MusicResult
		err     error
		wantErr bool
	}{
		{
			name: "Error fetching music not found for id zero",
			args: args{
				filter: entity.MusicFilter{
					ID: "0",
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
			name: "Success Get Music",
			args: args{
				filter: entity.MusicFilter{
					ID: "1",
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
			mockRepo := repository.NewMockMusicRepositoryItf(ctrl)
			mockRepo.EXPECT().GetMusic(gomock.Any()).Return(tt.want, tt.err)

			d := &Music{
				repo: mockRepo,
			}
			got, err := d.FetchMusic(&tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMusic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchMusic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicService_DetailMusic(t *testing.T) {
	type args struct {
		filter entity.MusicFilter
	}
	tests := []struct {
		name           string
		args           args
		expectedResult entity.MusicResult
		want           entity.Music
		err            error
		wantErr        bool
	}{
		{
			name: "Error get detail music not found",
			args: args{
				filter: entity.MusicFilter{
					ID: "0",
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
			name: "Success Get Detail Music",
			args: args{
				filter: entity.MusicFilter{
					ID: "1",
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
			expectedResult: entity.MusicResult{
				List: []entity.Music{
					{
						ID: 1,
					},
				},
			},
			want: entity.Music{
				ID: 1,
			},
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockMusicRepositoryItf(ctrl)
			mockRepo.EXPECT().GetMusic(gomock.Any()).Return(tt.expectedResult, tt.err)

			d := &Music{
				repo: mockRepo,
			}
			got, err := d.DetailMusic(&tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetailMusic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetailMusic() got = %v, want %v", got, tt.want)
			}
		})
	}
}
