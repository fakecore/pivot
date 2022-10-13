package sys

import (
	"github.com/gin-gonic/gin"
	"indulge/api/v1/sys"
	_ "indulge/api/v1/sys/impl"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type SysUserRouter struct {
	User sys.UserInterface `singleton:"indulge/api/v1/db/impl.User"`
}

func (r *SysUserRouter) InitRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("user")
	//apiRouterWithoutRecord := Router.Group("api")
	{
		apiRouter.POST("createUser", r.User.CreateUser) // 创建Api
	}
}
