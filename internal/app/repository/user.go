package repository

import (
	"fmt"
	"github.com/iss14036/music-chart/configs/schema"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/tools/driver"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	db driver.GormItf
}

func (u *User) IsUserLogin(filter *entity.UserFilter) error {
	panic("implement me")
}

func (u *User) GetUser(filter *entity.UserFilter) (entity.User, error) {
	var user entity.User

	q := u.db.Table("t_user").Select("*")
	q = u.constructFilter(filter, q)

	err := q.Find(&user).Error()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) constructFilter(filter *entity.UserFilter, q driver.GormItf) driver.GormItf {
	if filter.ID != "" {
		q = q.Where("id = ?", filter.ID)
	}

	if filter.Username != "" {
		q = q.Where("username = ?", filter.Username)
	}

	return q
}

func (u *User) InsertUser(user *entity.User) error {
	var err error

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	fmt.Println(user.Password)
	g := u.db.Begin()
	defer func() {
		if err != nil {
			g.Rollback()
		}
	}()

	err = g.Table("t_user").Save(&user).Error()
	if err != nil {
		return err
	}
	if err = g.Commit().Error(); err != nil {
		return err
	}

	return err
}

func (u *User) UpdateUser(user *entity.User) error {
	panic("implement me")
}

func NewUser(database schema.DatabaseReplication) UserRepositoryItf {
	return &User{
		db: database.Master,
	}
}
