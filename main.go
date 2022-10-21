package main

import (
	_ "github.com/go-spring/starter-echo"
	"povit/core"
)

func main() {
	core := core.Core{}
	core.Run()
}
