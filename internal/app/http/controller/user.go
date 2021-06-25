package controller

import (
	"github.com/iss14036/music-chart/internal/app/service"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/internal/pkg/responsewrapper"

	"github.com/labstack/echo/v4"
)

type UserCtrl struct {
	service service.UserServiceItf
}

func NewUser(service service.UserServiceItf) *UserCtrl {
	return &UserCtrl{
		service: service,
	}
}

func (u *UserCtrl) GetUser(c echo.Context) error {
	id := c.Param("id")
	filter := entity.UserFilter{
		ID: id,
	}

	res, err := u.service.GetUser(&filter)
	if err != nil {
		return responsewrapper.BadRequest(c, err, "")
	}

	return responsewrapper.OK(c, constant.MessageSuccessGet, res)
}

func (u *UserCtrl) InsertUser(c echo.Context) error {
	val := new(entity.User)
	err := c.Bind(val)
	if err != nil {
		return responsewrapper.BadRequest(c, err, "")
	}

	if err = val.Validate(); err != nil {
		return responsewrapper.BadRequest(c, err, "")
	}

	err = u.service.InsertUser(val)
	if err != nil {
		return responsewrapper.InternalServerError(c, err, "")
	}

	return responsewrapper.Created(c, constant.MessageSuccessCreate, "")
}

func (u *UserCtrl) Login(c echo.Context) error {
	val := new(entity.User)
	err := c.Bind(val)
	if err != nil {
		return responsewrapper.BadRequest(c, err, "")
	}
	filter := entity.UserFilter{
		Username: val.Username,
		Password: val.Password,
	}

	token, err := u.service.Login(&filter)
	if err != nil {
		return responsewrapper.InternalServerError(c, err, "")
	}

	data := entity.UserToken{Token: token}

	return responsewrapper.OK(c, constant.MessageSuccessLogin, data)
}
