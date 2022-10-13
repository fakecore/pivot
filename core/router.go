package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"indulge/router"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	private_group := Router.Group("")
	//TODO add middleware
	//private_group.Use()
	routerMain, err := router.GetRouterMainSingleton()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	//init router into gin
	routerMain.SysRouter.SysUser.InitRouter(private_group)

	return Router
}
