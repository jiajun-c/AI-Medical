package crud

import (
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func CheckTable(table interface{}) {
	exist, _ := Engine.IsTableExist(table)
	if !exist {
		err := Engine.CreateTables(table)
		if err != nil {
			log.Error(err)
			panic("failed to init the database")
		}
	}
}
