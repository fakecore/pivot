# pivot
A convenient gin-web backend service.
using IOC and reflection to enhance system init.

Thank @gin-vue-admin a lot, it's a great job. And it's inspiring me to in some structure and code design.

## How to use

It's a module-style code design. You can cut off the module useless.

## Environment

[ioc-golang](https://github.com/alibaba/ioc-golang)(we depend on its tool to generate compiled-time code)

## Project Key Point

We use IOC and Reflection to reduce the retrieved code.

see how it works.

ioc gets instance

```go
//about instance autowire

//get router instance
routerMain, err := router.GetRouterMainSingleton()


// +ioc:autowire=true
// +ioc:autowire:type=singleton
type RouterMain struct {
	SysRouter *sys.SysRouter `singleton:""`
}
//file sys_router.go
// +ioc:autowire=true
// +ioc:autowire:type=singleton
type SysRouter struct {
	SysUser RouterBase `singleton:"pivot/router/sys.SysUserRouter"`
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type SysUserRouter struct {
	User sys.UserInterface `singleton:"pivot/api/v1/sys/impl.User"`
}

type UserInterface interface {
	CreateUser(ctx *gin.Context)
}


// +ioc:autowire=true
// +ioc:autowire:type=singleton
type User struct {
}
func (u *User) CreateUser(ctx *gin.Context) {
}

```

reflection automatically register router function

```go
//routerMain.SysRouter.SysUser.InitRouter(private_group)
//equals 
	
	subRouterList := gsrf.GetStructFiledList(*routerMain)
	for _, c := range subRouterList {
		fi := gsrf.GetFieldInstanceByName(*routerMain, c)
		baseRouterList := gsrf.GetStructFieldListWithType(fi, "RouterBase")
		for _, sub := range baseRouterList {
			global.G_LOG.Info("baseRouter", zap.Any("", sub))
			tt := gsrf.GetFieldInstanceByName(fi, sub)
			gsrf.ExecMethod(tt, "InitRouter", private_group)
		}
	}
```

Think about if we have many router register functions. Then you need to write it `xx.xx.InitRouter` again and again, But now you have the auto-register utils. You only need to think about how to name your function.

## Demo



## Contribute

If you enjoy the code style of this project is. Please be generous with your ideas and have pr :)