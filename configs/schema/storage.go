package schema

import (
	"github.com/iss14036/music-chart/tools/driver"
)

// Database db replication configs struct
type DatabaseReplication struct {
	Master driver.GormItf
}
