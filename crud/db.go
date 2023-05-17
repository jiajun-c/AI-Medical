package crud

import "xorm.io/xorm"

var Engine *xorm.Engine

func CheckTable(table interface{}) {
	exist, _ := Engine.IsTableExist(table)
	if !exist {
		err := Engine.CreateTables(table)
		if err != nil {
			panic("failed to init the database")
		}
	}
}
