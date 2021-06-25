package controller

import (
	"strconv"
	"github.com/iss14036/music-chart/internal/app/service"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/internal/pkg/responsewrapper"

	"github.com/labstack/echo/v4"
)

type MusicCtrl struct {
	service service.MusicServiceItf
}

func NewMusic(service service.MusicServiceItf) *MusicCtrl {
	return &MusicCtrl{
		service: service,
	}
}

func (m *MusicCtrl) FetchMusic(c echo.Context) error {
	id := c.QueryParam("id")
	sortedField := c.QueryParam("sorted_field")
	sortedDirection := c.QueryParam("sorted_direction")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	row, _ := strconv.Atoi(c.QueryParam("row"))

	filter := entity.MusicFilter{
		ID: id,
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

	res, err := m.service.FetchMusic(&filter)
	if err != nil {
		return responsewrapper.BadRequest(c, err, "")
	}

	return responsewrapper.OK(c, constant.MessageSuccessGet, res)
}

func (m *MusicCtrl) DetailMusic(c echo.Context) error {
	id := c.Param("id")

	filter := entity.MusicFilter{
		ID: id,
		DataTable: entity.FilterDataTable{
			Pagination: entity.PaginationFilter{
				DisablePagination: true,
			},
		},
	}

	res, err := m.service.FetchMusic(&filter)
	if err != nil {
		return responsewrapper.BadRequest(c, err, "")
	}

	return responsewrapper.OK(c, constant.MessageSuccessGet, res)
}
