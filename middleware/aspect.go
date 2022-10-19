package middleware

import "github.com/gin-gonic/gin"

// before request truly is handled. Do something with current reqeust.

// Aspect
// 1.check user available
// 2.request url is matching his own role
// 3.record operator to db
func Aspect() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
