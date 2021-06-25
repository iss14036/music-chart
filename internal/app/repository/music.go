package repository

import (
	"fmt"
	"github.com/iss14036/music-chart/configs/schema"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/tools/driver"
)

type Music struct {
	db driver.GormItf
}

func (m *Music) GetMusic(filter *entity.MusicFilter) (entity.MusicResult, error) {
	var musics []entity.Music
	var count int

	if !filter.DataTable.IsValid() {
		return entity.MusicResult{}, constant.ErrorDataTableNotValid
	}

	q := m.db.Select("*").Table("t_music")

	err := q.Count(&count).Error()
	if err != nil {
		return entity.MusicResult{}, err
	}
	q = m.constructFilterMusic(filter, q)
	err = q.Find(&musics).Error()
	if err != nil {
		return entity.MusicResult{}, err
	}

	result := entity.MusicResult{
		List: musics,
		Pagination: entity.Pagination{
			Page:  filter.DataTable.Pagination.Page,
			Row:   filter.DataTable.Pagination.Limit,
			Count: count,
		},
	}
	return result, nil
}

func (m *Music) constructFilterMusic(filter *entity.MusicFilter, q driver.GormItf) driver.GormItf {
	orderQuery := fmt.Sprintf("%s %s", filter.DataTable.Sort.Field, filter.DataTable.Sort.Direction)

	if filter.ID != "" {
		q = q.Where("id = ?", filter.ID)
	}
	filter.DataTable.Pagination.SetOffset()

	if filter.DataTable.IsPaginated() {
		q = q.Order(orderQuery)
	}

	if filter.DataTable.IsPaginated() {
		q = q.Offset(filter.DataTable.Pagination.Offset).Limit(filter.DataTable.Pagination.Limit)
	}

	return q
}

func NewMusic(database schema.DatabaseReplication) MusicRepositoryItf {
	return &Music{
		db: database.Master,
	}
}
