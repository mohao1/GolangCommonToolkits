package Util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapEqual(t *testing.T) {
	//m1 := map[string]string{
	//	"张三": "1",
	//}
	//m2 := map[int]string{
	//	1: "1",
	//	2: "2",
	//}
	fmt.Println(reflect.DeepEqual(ts{name: "张三", age: 1}, ts2{age: 1, name: "张三"}))
}

type ts struct {
	name string
	age  int
}

type ts2 struct {
	age  int
	name string
}
