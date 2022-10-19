package core

import (
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

type Test1 struct {
}

func (t *Test1) Hello() {

}
func test() {

}
func TestAutoRegisterAllRouter(t *testing.T) {
	//autoRegisterAllRouter(test)
	autoRegisterAllRouter(Test1{})

}
