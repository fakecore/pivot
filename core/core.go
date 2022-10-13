package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"indulge/global"
	"time"
)

type server interface {
	ListenAndServe() error
}

type Core struct {
}

func (c *Core) init() {
	//global.G_CONF =
	v, conf := InitConfig()
	global.G_CONF = conf
	global.G_VIPER = v
	global.G_DB = Gorm() // init gorm connection pools
	global.G_Engine = InitRouter()
	fmt.Println("{}", global.G_CONF)
	//TODO redis

	if global.G_DB != nil {
		//create db
		db, _ := global.G_DB.DB()
		defer db.Close()
	} else {
	}
}

func (c *Core) initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func (c *Core) Run() {
	c.init()
	address := fmt.Sprintf(":%d", global.G_CONF.Sys.Port)
	fmt.Println(address)
	s := c.initServer(address, global.G_Engine)
	s.ListenAndServe()
}
