package Util

import (
	"fmt"
	"testing"
)

func TestMapEqual(t *testing.T) {
	fmt.Println(Equal(ts{name: "张三", age: 1}, ts{age: 1, name: "张三"}))
}

type ts struct {
	name string
	age  int
}

type ts2 struct {
	age  int
	name string
}
