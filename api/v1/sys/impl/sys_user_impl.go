package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type User struct {
}

func (u *User) CreateUser(ctx *gin.Context) {
	fmt.Println("hello hello,creating user now")
	//resp := db.HelloResp{}
	//return web.SUCCESS.Data(resp)
}
