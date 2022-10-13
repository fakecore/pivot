package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"indulge/global"
	"indulge/model/config"
)

func Gorm() *gorm.DB {
	//if global.G_CONF.Sys.Env == "prod"
	switch global.G_CONF.DB.Type {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func getMysqlConfigDns(db config.DBConfig) string {
	return db.UserName + ":" + db.Password + "@tcp(" + db.Url + ":" + db.Port + ")/" + db.DBName + "?" + db.Config
}
func GormMysql() *gorm.DB {
	dbc := global.G_CONF.DB
	if dbc.UserName == "" {
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       getMysqlConfigDns(dbc), // DSN data source name
		DefaultStringSize:         191,                    // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                  // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbc.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbc.MaxOpenConns)
		return db
	}
}
