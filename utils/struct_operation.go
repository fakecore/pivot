package utils

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/drgrib/iter"
	"reflect"
)

func StructName(src interface{}) (name string) {
	return reflect.TypeOf(src).Elem().Name()
}

// StructFunctionList must pass the struct pointer to get all the methods of src struct
func StructFunctionList(src interface{}) (fnList []list.List) {
	typeList := reflect.TypeOf(src)
	fmt.Println(typeList.NumMethod())
	for i := range iter.N(typeList.NumMethod()) {
		fmt.Println(typeList.Method(i).Name)
	}
	return nil
}

func StructPropertyList(src interface{}) (fnList []list.List) {
	return nil
}

//ExecMethod
//@param: src target struct.Should pass pointer type
//@param: fnName call function name
//@param: args function execute with those params
//fixme: check the count and type of args are matching target function
func ExecMethod(src interface{}, fnName string, args ...interface{}) (err error) {
	typeList := reflect.TypeOf(src)
	fn, isOK := typeList.MethodByName(fnName)
	if !isOK {
		err = errors.New(fmt.Sprintf("not found:%s in %s\n", fnName, typeList.Elem().Name()))
		return
	}
	var params []reflect.Value
	for _, c := range args {
		params = append(params, reflect.ValueOf(c))
	}
	fnCount := reflect.TypeOf(fn).NumMethod()
	if fnCount != len(params) {
		err = errors.New(fmt.Sprintf("param count error,target:%d,now:%d\n", fnCount, len(params)))
		return
	}
	reflect.ValueOf(src).MethodByName(fnName).Call(params)
	return nil
}

func CallFunc(src interface{}, fnName string) {
	reflect.ValueOf(&src).MethodByName(fnName).Call([]reflect.Value{})
}
