package global

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"indulge/model/config"
)

var (
	G_DB     *gorm.DB
	G_VIPER  *viper.Viper
	G_CONF   *config.Config
	G_Engine *gin.Engine
)
