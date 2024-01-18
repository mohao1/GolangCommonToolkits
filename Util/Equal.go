package Util

import "reflect"

/**
两种数据相同类型进行数据比较比较数据是否相同
*/

// Equal 相同类型数据比较
func Equal(m1, m2 any) bool {
	//nil情况下的处理逻辑
	if m1 == nil || m2 == nil {
		if m1 == nil && m2 == nil {
			return true
		}
		return false
	}

	v1 := reflect.ValueOf(m1)
	v2 := reflect.ValueOf(m2)
	if v1.Kind() != v2.Kind() {
		return false
	}

	switch v1.Kind() {
	case reflect.Map, reflect.Slice, reflect.Struct:
		return IncomparableType(m1, m2, v1.Kind())
	default:
		return v1 == v2
	}

}

func IncomparableType(m1, m2 any, kind reflect.Kind) bool {
	switch kind {
	case reflect.Slice, reflect.Map:
		reflect.DeepEqual(m1, m2)
	case reflect.Struct:
		if reflect.TypeOf(m1) == reflect.TypeOf(m2) {
			return reflect.DeepEqual(m1, m2)
		} else {
			return false
		}
	}
	return false
}
