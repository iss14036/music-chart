package service

import (
	"github.com/iss14036/music-chart/internal/app/repository"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/internal/pkg/gateway"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repository.UserRepositoryItf
	auth *gateway.Auth
}

func (u *User) Login(filter *entity.UserFilter) (token string, err error) {
	user, err := u.repo.GetUser(filter)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(filter.Password))
	if err != nil {
		return "", err
	}

	token, err = u.auth.GetToken(user.ID)
	return token, err
}

func (u *User) InsertUser(user *entity.User) error {
	return u.repo.InsertUser(user)
}

func (u *User) GetUser(filter *entity.UserFilter) (entity.User, error) {
	return u.repo.GetUser(filter)
}

func NewUser(repo repository.UserRepositoryItf, auth *gateway.Auth) UserServiceItf {
	return &User{
		repo: repo,
		auth: auth,
	}
}
