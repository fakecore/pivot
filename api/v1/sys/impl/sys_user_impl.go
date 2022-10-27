package impl

import (
	"github.com/gin-gonic/gin"
	"pivot/global"
	"pivot/model/db/biz"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type User struct {
}

func (u *User) CreateUser(ctx *gin.Context) {
	global.G_LOG.Info("hello hello,creating user now")
	//ctx.JSON(0, "{\"s\":\"hello hello,creating user now\"")
	item := []biz.Item{}
	global.G_DB.Find(&item)
	// fetch all places from the db

	//rows, err := global.G_DB.Queryx("SELECT * FROM items")
	//// iterate over each row
	//item := biz.Item{}
	//for rows.Next() {
	//	err = rows.StructScan(&item)
	//}
	//// check the error from rows
	//err = rows.Err()
	//global.G_LOG.Info("", zap.Any("", err))
	////global.G_LOG.Info("", zap.Any("", )
	//global.G_LOG.Info("", zap.Any("", &item))

	//resp := db.HelloResp{}
	//return web.SUCCESS.Data(resp)
}
