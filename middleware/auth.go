package middleware

import (
	"github.com/gin-gonic/gin"
	"pivot/global"
	"pivot/model/domain"
	"pivot/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			domain.FailCodeStr(c, global.ERROR_TOKEN_EMPTY)
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		_, err := utils.NewJWT().ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				domain.FailCodeStr(c, global.ERROR_TOKEN_EXPIRED)
			} else {
				domain.FailCode(c, global.ERROR_FAILED, err.Error())
			}
			c.Abort()
			return
		}
		c.Next()
	}
}
