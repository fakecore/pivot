package sys

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type SysRouter struct {
	SysUser RouterBase `singleton:"povit/router/sys.SysUserRouter"`
}
