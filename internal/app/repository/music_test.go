package repository

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/tools/driver"
	"reflect"
	"testing"
)

func TestMusic_GetMusic(t *testing.T) {
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
			name: "Error when music not found",
			args: args{
				filter: entity.MusicFilter{
					ID: "1",
					DataTable: entity.FilterDataTable{
						Pagination: entity.PaginationFilter{
							DisablePagination: true,
						},
					},
				},
			},
			err:     errors.New("record not found"),
			wantErr: true,
		},
		{
			name: "Success get musics",
			args: args{
				filter: entity.MusicFilter{
					ID:     "1",
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
			want: entity.MusicResult{
				List: []entity.Music{
					entity.Music{
						ID:          1,
						Title:       "test",
						Singer:      "test",
						Duration:    "123",
						Album:       "test",
						ReleaseYear: "2004",
					},
				},
				Pagination: entity.Pagination{
					Page:  1,
					Row:   10,
					Count: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			var musics []entity.Music
			var count int
			gorm := driver.NewMockGormItf(ctrl)
			gorm.EXPECT().Table("t_music").Return(gorm).AnyTimes()
			gorm.EXPECT().Select("*").Return(gorm).AnyTimes()
			gorm.EXPECT().Where("id = ?", gomock.Any()).Return(gorm).AnyTimes()
			gorm.EXPECT().Count(&count).Return(gorm).AnyTimes()
			gorm.EXPECT().Offset(gomock.Any()).Return(gorm).AnyTimes()
			gorm.EXPECT().Limit(gomock.Any()).Return(gorm).AnyTimes()
			gorm.EXPECT().Order(gomock.Any()).Return(gorm).AnyTimes()
			if len(tt.want.List) != 0 {
				gorm.EXPECT().Find(&musics).
					Do(func(p *[]entity.Music) {
						for _, item := range tt.want.List {
							*p = append(*p, item)
						}
					}).Return(gorm)
			}
			gorm.EXPECT().Error().Return(tt.err).AnyTimes()

			d := &Music{
				db: gorm,
			}
			got, err := d.GetMusic(&tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMusic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMusic() got = %v, want %v", got, tt.want)
			}
		})
	}
}
