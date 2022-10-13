package sys

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type SysRouter struct {
	SysUser RouterBase `singleton:"indulge/router/db.SysUserRouter"`
}
