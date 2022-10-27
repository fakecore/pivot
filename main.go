package main

import (
	_ "github.com/go-spring/starter-echo"
	"pivot/core"
)

func main() {
	core := core.Core{}
	core.Run()
}
