package sys

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type SysRouter struct {
	SysUser RouterBase `singleton:"pivot/router/sys.SysUserRouter"`
}
