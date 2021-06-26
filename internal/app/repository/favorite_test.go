package repository

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/tools/driver"
	"reflect"
	"testing"
)

func TestFavorite_SetFavorite(t *testing.T) {
	type args struct {
		sg entity.Favorite
	}
	tests := []struct {
		name      string
		args      args
		errSave   error
		wantErr   bool
	}{
		{
			name: "Success set favorite",
			args: args{
				sg: entity.Favorite{
					UserID:  1,
					MusicID: 1,
				},
			},
		},
		{
			name: "Failed set favorite",
			args: args{
				sg: entity.Favorite{},
			},
			errSave:   errors.New("error saved"),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockGorm := driver.NewMockGormItf(ctrl)
			mockGorm.EXPECT().Table("t_favorite").Return(mockGorm).AnyTimes()
			mockGorm.EXPECT().Save(gomock.Any()).Return(mockGorm).AnyTimes()
			mockGorm.EXPECT().Error().Return(tt.errSave)

			d := &Favorite{
				db: mockGorm,
			}
			err := d.SetFavorite(&tt.args.sg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetFavorite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFavorite_GetFavorite(t *testing.T) {
	type args struct {
		filter entity.FavoriteFilter
	}
	tests := []struct {
		name    string
		args    args
		want    entity.FavoriteResult
		err     error
		wantErr bool
	}{
		{
			name: "Error when favorite not found",
			args: args{
				filter: entity.FavoriteFilter{
					UserID: "0",
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
			name: "Success get Favorite Music",
			args: args{
				filter: entity.FavoriteFilter{
					UserID:     "1",
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
			want: entity.FavoriteResult{
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
			gorm.EXPECT().Table("t_favorite").Return(gorm).AnyTimes()
			gorm.EXPECT().Table("t_music tm").Return(gorm).AnyTimes()
			gorm.EXPECT().Select("*").Return(gorm).AnyTimes()
			gorm.EXPECT().Select("tm.*").Return(gorm).AnyTimes()
			gorm.EXPECT().Joins("INNER JOIN t_favorite tf on tf.music_id = tm.id").Return(gorm).AnyTimes()
			gorm.EXPECT().Where("tf.user_id = ?", gomock.Any()).Return(gorm).AnyTimes()
			gorm.EXPECT().Where("user_id = ?", gomock.Any()).Return(gorm).AnyTimes()
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

			d := &Favorite{
				db: gorm,
			}
			got, err := d.GetFavorite(&tt.args.filter)
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
