package service

import (
	"github.com/iss14036/music-chart/internal/app/repository"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"strconv"
)

type Favorite struct {
	repo      repository.FavoriteRepositoryItf
	musicRepo repository.MusicRepositoryItf
}

func (f *Favorite) SetFavorite(favorite *entity.Favorite) error {
	musicFilter := &entity.MusicFilter{
		ID: strconv.Itoa(favorite.MusicID),
		DataTable: entity.FilterDataTable{
			Pagination: entity.PaginationFilter{
				DisablePagination: true,
			},
		},
	}
	_, err := f.musicRepo.GetMusic(musicFilter)
	if err != nil {
		return constant.ErrorMusicNotFound
	}

	return f.repo.SetFavorite(favorite)
}

func (f *Favorite) GetFavorite(filter *entity.FavoriteFilter) (entity.FavoriteResult, error) {
	return f.repo.GetFavorite(filter)
}

func NewFavorite(repo repository.FavoriteRepositoryItf, musicRepo repository.MusicRepositoryItf) FavoriteServiceItf {
	return &Favorite{
		repo:      repo,
		musicRepo: musicRepo,
	}
}
