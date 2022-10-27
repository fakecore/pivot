package core

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pivot/global"
	"pivot/model/config"
)

func InitDbOrDie() *gorm.DB {
	//if global.G_CONF.Sys.Env == "prod"
	switch global.G_CONF.DB.Type {
	case "mysql":
		return GormMysqlOrDie()
	default:
		return GormMysqlOrDie()
	}
}

func GormMysqlOrDie() *gorm.DB {
	//
	dbc := global.G_CONF.DB.Mysql
	if dbc.Url == "" {
		panic("the db's url is empty")
	}
	db, err := gorm.Open(mysql.Open(getMysqlConfigDns(dbc)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func getMysqlConfigDns(db config.MysqlConfig) string {
	connInfo := db.UserName + ":" + db.Password + "@tcp(" + db.Url + ")/" + db.DbName + "?" + db.Config
	global.G_LOG.Info("", zap.String("connecting: ", connInfo))
	return db.UserName + ":" + db.Password + "@tcp(" + db.Url + ")/" + db.DbName + "?" + db.Config
}
