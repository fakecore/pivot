package utils

import (
	"fmt"
	"testing"
)

type TestF struct {
}

func (b TestF) HiV() {
	fmt.Println("hiv")
}

func (b *TestF) Hi(s string) {
	fmt.Println("hi,this is hi", s)
}

func TestOperation(t *testing.T) {
	//StructFunctionList(&TestF{})
	err := ExecMethod(&TestF{}, "Hi", "eeeee", "ddddd")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ExecMethod(&TestF{}, "HiV", "eeeee", "ddddd")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ExecMethod(&TestF{}, "NooooV", "eeeee", "ddddd")
	if err != nil {
		fmt.Println(err.Error())
	}
}
