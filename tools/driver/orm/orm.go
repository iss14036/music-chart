package orm

import (
	"database/sql"
	"github.com/iss14036/music-chart/tools/driver"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type gormDB = gorm.DB

type gormw struct {
	*gormDB
}

// Openw is a drop-in replacement for Open()
func Openw(dialect string, args ...interface{}) (db driver.GormItf, err error) {
	gormdb, err := gorm.Open(dialect, args...)
	return Wrap(gormdb), err
}

// Wrap wraps gorm.DB in an interface
func Wrap(db *gorm.DB) driver.GormItf {
	return &gormw{db}
}

func (it *gormw) Close() error {
	return it.gormDB.Close()
}

func (it *gormw) DB() *sql.DB {
	return it.gormDB.DB()
}

func (it *gormw) New() driver.GormItf {
	return Wrap(it.gormDB.New())
}

func (it *gormw) NewScope(value interface{}) *gorm.Scope {
	return it.gormDB.NewScope(value)
}

func (it *gormw) CommonDB() gorm.SQLCommon {
	return it.gormDB.CommonDB()
}

func (it *gormw) Callback() *gorm.Callback {
	return it.gormDB.Callback()
}

func (it *gormw) SetLogger(log gorm.Logger) {
	it.gormDB.SetLogger(log)
}

func (it *gormw) LogMode(enable bool) driver.GormItf {
	return Wrap(it.gormDB.LogMode(enable))
}

func (it *gormw) SingularTable(enable bool) {
	it.gormDB.SingularTable(enable)
}

func (it *gormw) Where(query interface{}, args ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Where(query, args...))
}

func (it *gormw) Or(query interface{}, args ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Or(query, args...))
}

func (it *gormw) Not(query interface{}, args ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Not(query, args...))
}

func (it *gormw) Limit(value int) driver.GormItf {
	return Wrap(it.gormDB.Limit(value))
}

func (it *gormw) Offset(value int) driver.GormItf {
	return Wrap(it.gormDB.Offset(value))
}

func (it *gormw) Order(value string, reorder ...bool) driver.GormItf {
	return Wrap(it.gormDB.Order(value, reorder...))
}

func (it *gormw) Select(query interface{}, args ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Select(query, args...))
}

func (it *gormw) Omit(columns ...string) driver.GormItf {
	return Wrap(it.gormDB.Omit(columns...))
}

func (it *gormw) Group(query string) driver.GormItf {
	return Wrap(it.gormDB.Group(query))
}

func (it *gormw) Having(query string, values ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Having(query, values...))
}

func (it *gormw) Joins(query string, args ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Joins(query, args...))
}

func (it *gormw) Scopes(funcs ...func(*gorm.DB) *gorm.DB) driver.GormItf {
	return Wrap(it.gormDB.Scopes(funcs...))
}

func (it *gormw) Unscoped() driver.GormItf {
	return Wrap(it.gormDB.Unscoped())
}

func (it *gormw) Attrs(attrs ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Attrs(attrs...))
}

func (it *gormw) Assign(attrs ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Assign(attrs...))
}

func (it *gormw) First(out interface{}, where ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.First(out, where...))
}

func (it *gormw) Last(out interface{}, where ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Last(out, where...))
}

func (it *gormw) Find(out interface{}, where ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Find(out, where...))
}

func (it *gormw) Scan(dest interface{}) driver.GormItf {
	return Wrap(it.gormDB.Scan(dest))
}

func (it *gormw) Row() *sql.Row {
	return it.gormDB.Row()
}

func (it *gormw) Rows() (*sql.Rows, error) {
	return it.gormDB.Rows()
}

func (it *gormw) ScanRows(rows *sql.Rows, result interface{}) error {
	return it.gormDB.ScanRows(rows, result)
}

func (it *gormw) Pluck(column string, value interface{}) driver.GormItf {
	return Wrap(it.gormDB.Pluck(column, value))
}

func (it *gormw) Count(value interface{}) driver.GormItf {
	return Wrap(it.gormDB.Count(value))
}

func (it *gormw) Related(value interface{}, foreignKeys ...string) driver.GormItf {
	return Wrap(it.gormDB.Related(value, foreignKeys...))
}

func (it *gormw) FirstOrInit(out interface{}, where ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.FirstOrInit(out, where...))
}

func (it *gormw) FirstOrCreate(out interface{}, where ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.FirstOrCreate(out, where...))
}

func (it *gormw) Update(attrs ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Update(attrs...))
}

func (it *gormw) Updates(values interface{}, ignoreProtectedAttrs ...bool) driver.GormItf {
	return Wrap(it.gormDB.Updates(values, ignoreProtectedAttrs...))
}

func (it *gormw) UpdateColumn(attrs ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.UpdateColumn(attrs...))
}

func (it *gormw) UpdateColumns(values interface{}) driver.GormItf {
	return Wrap(it.gormDB.UpdateColumns(values))
}

func (it *gormw) Save(value interface{}) driver.GormItf {
	return Wrap(it.gormDB.Save(value))
}

func (it *gormw) Create(value interface{}) driver.GormItf {
	return Wrap(it.gormDB.Create(value))
}

func (it *gormw) Delete(value interface{}, where ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Delete(value, where...))
}

func (it *gormw) Raw(sql string, values ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Raw(sql, values...))
}

func (it *gormw) Exec(sql string, values ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Exec(sql, values...))
}

func (it *gormw) Model(value interface{}) driver.GormItf {
	return Wrap(it.gormDB.Model(value))
}

func (it *gormw) Table(name string) driver.GormItf {
	return Wrap(it.gormDB.Table(name))
}

func (it *gormw) Debug() driver.GormItf {
	return Wrap(it.gormDB.Debug())
}

func (it *gormw) Begin() driver.GormItf {
	return Wrap(it.gormDB.Begin())
}

func (it *gormw) Commit() driver.GormItf {
	return Wrap(it.gormDB.Commit())
}

func (it *gormw) Rollback() driver.GormItf {
	return Wrap(it.gormDB.Rollback())
}

func (it *gormw) RollbackUnlessCommitted() driver.GormItf {
	return Wrap(it.gormDB.RollbackUnlessCommitted())
}

func (it *gormw) NewRecord(value interface{}) bool {
	return it.gormDB.NewRecord(value)
}

func (it *gormw) RecordNotFound() bool {
	return it.gormDB.RecordNotFound()
}

func (it *gormw) CreateTable(values ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.CreateTable(values...))
}

func (it *gormw) DropTable(values ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.DropTable(values...))
}

func (it *gormw) DropTableIfExists(values ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.DropTableIfExists(values...))
}

func (it *gormw) HasTable(value interface{}) bool {
	return it.gormDB.HasTable(value)
}

func (it *gormw) AutoMigrate(values ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.AutoMigrate(values...))
}

func (it *gormw) ModifyColumn(column string, typ string) driver.GormItf {
	return Wrap(it.gormDB.ModifyColumn(column, typ))
}

func (it *gormw) DropColumn(column string) driver.GormItf {
	return Wrap(it.gormDB.DropColumn(column))
}

func (it *gormw) AddIndex(indexName string, columns ...string) driver.GormItf {
	return Wrap(it.gormDB.AddIndex(indexName, columns...))
}

func (it *gormw) AddUniqueIndex(indexName string, columns ...string) driver.GormItf {
	return Wrap(it.gormDB.AddUniqueIndex(indexName, columns...))
}

func (it *gormw) RemoveIndex(indexName string) driver.GormItf {
	return Wrap(it.gormDB.RemoveIndex(indexName))
}

func (it *gormw) Association(column string) *gorm.Association {
	return it.gormDB.Association(column)
}

func (it *gormw) Preload(column string, conditions ...interface{}) driver.GormItf {
	return Wrap(it.gormDB.Preload(column, conditions...))
}

func (it *gormw) Set(name string, value interface{}) driver.GormItf {
	return Wrap(it.gormDB.Set(name, value))
}

func (it *gormw) InstantSet(name string, value interface{}) driver.GormItf {
	return Wrap(it.gormDB.InstantSet(name, value))
}

func (it *gormw) Get(name string) (interface{}, bool) {
	return it.gormDB.Get(name)
}

func (it *gormw) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	it.gormDB.SetJoinTableHandler(source, column, handler)
}

func (it *gormw) AddForeignKey(field string, dest string, onDelete string, onUpdate string) driver.GormItf {
	return Wrap(it.gormDB.AddForeignKey(field, dest, onDelete, onUpdate))
}

func (it *gormw) AddError(err error) error {
	return it.gormDB.AddError(err)
}

func (it *gormw) GetErrors() (errors []error) {
	return it.gormDB.GetErrors()
}

func (it *gormw) RowsAffected() int64 {
	return it.gormDB.RowsAffected
}

func (it *gormw) Error() error {
	return it.gormDB.Error
}
