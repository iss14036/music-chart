package controller

import (
	"fmt"
	"strconv"
	"github.com/iss14036/music-chart/internal/app/service"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/internal/pkg/responsewrapper"

	"github.com/labstack/echo/v4"
)

type FavoriteCtrl struct {
	service service.FavoriteServiceItf
}

func NewFavorite(service service.FavoriteServiceItf) *FavoriteCtrl {
	return &FavoriteCtrl{
		service: service,
	}
}

func (f *FavoriteCtrl) SetFavorite(c echo.Context) error {
	val := new(entity.Favorite)
	err := c.Bind(val)
	if err != nil {
		return responsewrapper.BadRequest(c, err, "")
	}
	val.UserID = int(c.Get("user_id").(float64))

	err = f.service.SetFavorite(val)
	if err != nil {
		return responsewrapper.InternalServerError(c, err, "")
	}

	return responsewrapper.Created(c, constant.MessageSuccessCreate, "")
}

func (f *FavoriteCtrl) FetchFavorite(c echo.Context) error {
	userID := c.Get("user_id")
	sortedField := c.QueryParam("sorted_field")
	sortedDirection := c.QueryParam("sorted_direction")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	row, _ := strconv.Atoi(c.QueryParam("row"))

	filter := entity.FavoriteFilter{
		UserID: fmt.Sprintf("%v", userID),
		DataTable: entity.FilterDataTable{
			Sort: entity.Sort{
				Field:     sortedField,
				Direction: sortedDirection,
			},
			Pagination: entity.PaginationFilter{
				Page:  page,
				Limit: row,
			},
		},
	}

	res, err := f.service.GetFavorite(&filter)
	if err != nil {
		return responsewrapper.InternalServerError(c, err, "")
	}

	return responsewrapper.OK(c, constant.MessageSuccessGet, res)
}
