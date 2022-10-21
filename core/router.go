package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"povit/middleware"
	"povit/router"
	"reflect"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	//public_group := Router.Group("")
	private_group := Router.Group("")
	//private_group.Use()

	//TODO add middleware
	private_group.Use(middleware.JWTAuth())
	routerMain, err := router.GetRouterMainSingleton()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	//init router into gin
	routerMain.SysRouter.SysUser.InitRouter(private_group)

	return Router
}

func autoRegisterAllRouter(ptr interface{}) {
	//fnName := runtime.FuncForPC(reflect.ValueOf(ptr).Pointer()).Name()
	//fmt.Println(fnName)

	fnType, fnValue := reflect.TypeOf(ptr), reflect.ValueOf(ptr)
	propertyNums := fnType.NumField()
	for i := 0; i < propertyNums; i++ {
		// 属性
		property := fnType.Field(i)
		// 待填充属性值
		propertyValue := fnValue.FieldByName(property.Name)
		fmt.Printf("%v %v", propertyValue, property)
	}

	//fields := strings.FieldsFunc(fnName, func(sep rune) bool {
	//	for _, s := range seps {
	//		if sep == s {
	//			return true
	//		}
	//	}
	//	return false
	//})

}
