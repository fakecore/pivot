package db

import (
	"github.com/casbin/gorm-adapter/v3"
)

var carbin_rules = []gormadapter.CasbinRule{
	{Ptype: "p", V0: "1", V1: "/base/login", V2: "POST"},
}
