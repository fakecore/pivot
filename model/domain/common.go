package domain

import (
	"github.com/gin-gonic/gin"
	"indulge/global"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Fail(c *gin.Context, str string) {
	c.JSON(http.StatusOK, Response{
		Msg: str,
	})
}

func FailCode(c *gin.Context, code int, str string) {
	c.JSON(http.StatusOK, Response{
		Msg:  str,
		Code: code,
	})
}

func FailCodeStr(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Msg:  global.ERROR_STRING[code],
		Code: code,
	})
}
