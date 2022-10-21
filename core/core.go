package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"povit/global"
	"time"
)

type server interface {
	ListenAndServe() error
}

type Core struct {
}

func (c *Core) init() {
	//global.G_CONF =
	//TODO using tags and reflection to automatically inject component
	v, conf := InitConfig()
	global.G_LOG = InitLogger()
	global.G_CONF = conf
	global.G_VIPER = v
	global.G_DB = InitDbOrDie()       // init gorm connection pools
	global.G_REDIS = InitRedisOrDie() // init gorm connection pools
	global.G_Engine = InitRouter()
	//global.G_LOG
	global.G_LOG.Info("hello, Welcome to povit!")
	//global.G_LOG.Info("%v", zap.Any("", global.G_CONF))
	//fmt.Println("{}", global.G_CONF)
	//TODO redis

	//if global.G_DB != nil {
	//create db
	//db, _ := global.G_DB.DB.
	//defer db.Close()
	//} else {
	//}
}

func (c *Core) deconstruct() {
	global.G_LOG.Sync()
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
	defer c.deconstruct()
	address := fmt.Sprintf(":%d", global.G_CONF.Sys.Port)
	fmt.Println(address)
	s := c.initServer(address, global.G_Engine)
	s.ListenAndServe()

}
