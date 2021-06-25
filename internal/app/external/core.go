package external

import (
	"log"
	"github.com/iss14036/music-chart/configs"
	"github.com/iss14036/music-chart/configs/schema"
	"github.com/iss14036/music-chart/tools/driver/orm"
)

func ProviderOrm(config *configs.Config) schema.DatabaseReplication {
	dbDriver := config.DbDriver
	masterConn := config.DBUri

	master, err := orm.Openw(dbDriver, masterConn)
	if err != nil {
		log.Fatalf("Cannot open master database: %s", err)
	}
	master.SingularTable(true)
	master.LogMode(config.DbLog)

	return schema.DatabaseReplication{
		Master: master,
	}
}
