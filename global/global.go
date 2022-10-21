package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/golang/groupcache/singleflight"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"povit/model/config"
)

var (
	G_DB               *gorm.DB
	G_VIPER            *viper.Viper
	G_CONF             *config.Config
	G_Engine           *gin.Engine
	G_LOG              *zap.Logger
	G_REDIS            *redis.Client
	G_Concurrency_Work = &singleflight.Group{}
)

func DEBUG(str *string) {
	G_LOG.Debug(*str)
}

func INFO(str *string) {
	G_LOG.Info(*str)
}

func WARN(str *string) {
	G_LOG.Warn(*str)
}

func ERROR(str *string) {
	G_LOG.Error(*str)
}

func FATAL(str *string) {
	G_LOG.Fatal(*str)
}
