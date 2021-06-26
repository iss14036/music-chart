package driver

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

//go:generate mockgen -destination=./driver_mock.go -package=driver -source=./driver.go
// GormItf is an interface which DB implements
type GormItf interface {
	Close() error
	DB() *sql.DB
	New() GormItf
	NewScope(value interface{}) *gorm.Scope
	CommonDB() gorm.SQLCommon
	Callback() *gorm.Callback
	SetLogger(l gorm.Logger)
	LogMode(enable bool) GormItf
	SingularTable(enable bool)
	Where(query interface{}, args ...interface{}) GormItf
	Or(query interface{}, args ...interface{}) GormItf
	Not(query interface{}, args ...interface{}) GormItf
	Limit(value int) GormItf
	Offset(value int) GormItf
	Order(value string, reorder ...bool) GormItf
	Select(query interface{}, args ...interface{}) GormItf
	Omit(columns ...string) GormItf
	Group(query string) GormItf
	Having(query string, values ...interface{}) GormItf
	Joins(query string, args ...interface{}) GormItf
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) GormItf
	Unscoped() GormItf
	Attrs(attrs ...interface{}) GormItf
	Assign(attrs ...interface{}) GormItf
	First(out interface{}, where ...interface{}) GormItf
	Last(out interface{}, where ...interface{}) GormItf
	Find(out interface{}, where ...interface{}) GormItf
	Scan(dest interface{}) GormItf
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) GormItf
	Count(value interface{}) GormItf
	Related(value interface{}, foreignKeys ...string) GormItf
	FirstOrInit(out interface{}, where ...interface{}) GormItf
	FirstOrCreate(out interface{}, where ...interface{}) GormItf
	Update(attrs ...interface{}) GormItf
	Updates(values interface{}, ignoreProtectedAttrs ...bool) GormItf
	UpdateColumn(attrs ...interface{}) GormItf
	UpdateColumns(values interface{}) GormItf
	Save(value interface{}) GormItf
	Create(value interface{}) GormItf
	Delete(value interface{}, where ...interface{}) GormItf
	Raw(sql string, values ...interface{}) GormItf
	Exec(sql string, values ...interface{}) GormItf
	Model(value interface{}) GormItf
	Table(name string) GormItf
	Debug() GormItf
	Begin() GormItf
	Commit() GormItf
	Rollback() GormItf
	RollbackUnlessCommitted() GormItf
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	CreateTable(values ...interface{}) GormItf
	DropTable(values ...interface{}) GormItf
	DropTableIfExists(values ...interface{}) GormItf
	HasTable(value interface{}) bool
	AutoMigrate(values ...interface{}) GormItf
	ModifyColumn(column string, typ string) GormItf
	DropColumn(column string) GormItf
	AddIndex(indexName string, column ...string) GormItf
	AddUniqueIndex(indexName string, column ...string) GormItf
	RemoveIndex(indexName string) GormItf
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) GormItf
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) GormItf
	Set(name string, value interface{}) GormItf
	InstantSet(name string, value interface{}) GormItf
	Get(name string) (value interface{}, ok bool)
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	AddError(err error) error
	GetErrors() (errors []error)

	// extra
	Error() error
	RowsAffected() int64
}
