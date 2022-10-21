package utils

import (
	"fmt"
	"testing"
)

type SrcData struct {
	v1   int
	Test string
}
type DestData struct {
	Uk   int
	v1   int
	Test string
}

func TestStructCopy(t *testing.T) {
	s := SrcData{v1: 23, Test: "hhhhh"}
	var d DestData
	StructCopy(&s, &d)
	fmt.Printf("%v %v", s, d)
	if s.v1 == d.v1 || s.Test != d.Test {
		t.Error("copy failed")
	}
}
