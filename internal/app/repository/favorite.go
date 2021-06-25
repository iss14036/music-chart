package repository

import (
	"fmt"
	"github.com/iss14036/music-chart/configs/schema"
	"github.com/iss14036/music-chart/internal/pkg/constant"
	"github.com/iss14036/music-chart/internal/pkg/entity"
	"github.com/iss14036/music-chart/tools/driver"
)

type Favorite struct {
	db driver.GormItf
}

func (f *Favorite) SetFavorite(favorite *entity.Favorite) error {
	err := f.db.Table("t_favorite").Save(&favorite).Error()

	return err
}

func (f *Favorite) GetFavorite(filter *entity.FavoriteFilter) (entity.FavoriteResult, error) {
	var favorites []entity.Music
	var count int

	if !filter.DataTable.IsValid() {
		return entity.FavoriteResult{}, constant.ErrorDataTableNotValid
	}

	q := f.db.Select("*").Table("t_favorite").Where("user_id = ?", filter.UserID)

	err := q.Count(&count).Error()
	if err != nil {
		return entity.FavoriteResult{}, err
	}
	q = f.db.Select("tm.*").Table("t_music tm").
		Joins("INNER JOIN t_favorite tf on tf.music_id = tm.id").
		Where("tf.user_id = ?", filter.UserID)
	q = f.constructFilterFavorite(filter, q)
	err = q.Find(&favorites).Error()
	if err != nil {
		return entity.FavoriteResult{}, err
	}

	return entity.FavoriteResult{
		List:       favorites,
		Pagination: entity.Pagination{
			Page:  filter.DataTable.Pagination.Page,
			Row:   filter.DataTable.Pagination.Limit,
			Count: count,
		},
	}, nil
}

func (f *Favorite) constructFilterFavorite(filter *entity.FavoriteFilter, q driver.GormItf) driver.GormItf {
	orderQuery := fmt.Sprintf("%s %s", filter.DataTable.Sort.Field, filter.DataTable.Sort.Direction)
	filter.DataTable.Pagination.SetOffset()

	if filter.DataTable.IsPaginated() {
		q = q.Order(orderQuery)
	}

	if filter.DataTable.IsPaginated() {
		q = q.Offset(filter.DataTable.Pagination.Offset).Limit(filter.DataTable.Pagination.Limit)
	}

	return q
}

func NewFavorite(database schema.DatabaseReplication) FavoriteRepositoryItf {
	return &Favorite{
		db: database.Master,
	}
}
