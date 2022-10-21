package sys

import (
	"github.com/gin-gonic/gin"
	"povit/api/v1/sys"
	_ "povit/api/v1/sys/impl"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type SysUserRouter struct {
	User sys.UserInterface `singleton:"povit/api/v1/sys/impl.User"`
}

func (r *SysUserRouter) InitRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("user")
	//apiRouterWithoutRecord := Router.Group("api")
	{
		apiRouter.POST("createUser", r.User.CreateUser) // 创建Api
	}
}
