package repository

import "github.com/iss14036/music-chart/internal/pkg/entity"

//go:generate mockgen -destination=./repository_mock.go -package=repository -source=./repository.go
type UserRepositoryItf interface {
	GetUser(filter *entity.UserFilter) (entity.User, error)
	InsertUser(user *entity.User) error
}

type MusicRepositoryItf interface {
	GetMusic(filter *entity.MusicFilter) (entity.MusicResult, error)
}

type FavoriteRepositoryItf interface {
	SetFavorite(favorite *entity.Favorite) error
	GetFavorite(filter *entity.FavoriteFilter) (entity.FavoriteResult, error)
}
