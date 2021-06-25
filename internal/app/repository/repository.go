package repository

import "github.com/iss14036/music-chart/internal/pkg/entity"

type UserRepositoryItf interface {
	GetUser(filter *entity.UserFilter) (entity.User, error)
	IsUserLogin(filter *entity.UserFilter) error
	InsertUser(user *entity.User) error
	UpdateUser(user *entity.User) error
}

type MusicRepositoryItf interface {
	GetMusic(filter *entity.MusicFilter) (entity.MusicResult, error)
}

type FavoriteRepositoryItf interface {
	SetFavorite(favorite *entity.Favorite) error
	GetFavorite(filter *entity.FavoriteFilter) (entity.FavoriteResult, error)
}
