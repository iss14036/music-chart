package service

import "github.com/iss14036/music-chart/internal/pkg/entity"

type UserServiceItf interface {
	GetUser(filter *entity.UserFilter) (entity.User, error)
	InsertUser(user *entity.User) error
	Login(filter *entity.UserFilter) (token string, err error)
}

type MusicServiceItf interface {
	FetchMusic(filter *entity.MusicFilter) (entity.MusicResult, error)
	DetailMusic(filter *entity.MusicFilter) (entity.Music, error)
}

type FavoriteServiceItf interface {
	SetFavorite(favorite *entity.Favorite) error
	GetFavorite(filter *entity.FavoriteFilter) (entity.FavoriteResult, error)
}
