package router

import (
	"indulge/router/sys"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type RouterMain struct {
	SysRouter *sys.SysRouter `singleton:""`
}
