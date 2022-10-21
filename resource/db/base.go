package db

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"povit/global"
	"povit/utils"
)

func InitTableOrDie(db *gorm.DB, dbTable interface{}) {
	global.G_LOG.Info("InitTable-" + utils.StructName(dbTable))
	err := global.G_DB.AutoMigrate(dbTable)
	if err != nil {
		panic(err)
	}
}

func InitDataOrDie(db *gorm.DB, dbTable interface{}) {
	global.G_LOG.Info("InitData-" + utils.StructName(gormadapter.CasbinRule{}))
	if db.Model(&dbTable).RowsAffected != 0 {
		errStr := fmt.Sprintf("DB:%vs already exists data,please empty the former data before initialization.", utils.StructName(dbTable))
		panic(errStr)
	}
	err := db.Create(&carbin_rules)
	if err != nil {
		panic(err)
	}
}
