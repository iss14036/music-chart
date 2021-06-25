package service

import (
	"github.com/iss14036/music-chart/internal/app/repository"
	"github.com/iss14036/music-chart/internal/pkg/entity"
)

type Music struct {
	repo repository.MusicRepositoryItf
}

func (m *Music) FetchMusic(filter *entity.MusicFilter) (entity.MusicResult, error) {
	return m.repo.GetMusic(filter)
}

func (m *Music) DetailMusic(filter *entity.MusicFilter) (entity.Music, error) {
	musicResult, err := m.repo.GetMusic(filter)
	if err != nil {
		return entity.Music{}, err
	}

	return musicResult.List[0], nil
}

func NewMusic(repo repository.MusicRepositoryItf) MusicServiceItf {
	return &Music{
		repo: repo,
	}
}
