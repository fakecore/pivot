package sys

import "github.com/gin-gonic/gin"

type RouterBase interface {
	InitRouter(Router *gin.RouterGroup)
}
