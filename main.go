package main

import (
	_ "github.com/go-spring/starter-echo"
	"indulge/core"
)

func main() {
	core := core.Core{}
	core.Run()
}
