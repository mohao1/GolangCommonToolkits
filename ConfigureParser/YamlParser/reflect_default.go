package YamlParser

import (
	"fmt"
	"reflect"
)

// setDefaults 使用反射设置默认值
func setDefaults(v reflect.Value) {
	t := v.Type()

	// 遍历结构体的每个字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		structField := t.Field(i)

		// 跳过不可设置的字段
		if !field.CanSet() {
			continue
		}

		// 处理嵌套结构体
		if field.Kind() == reflect.Struct {
			setDefaults(field)
			continue
		}

		// 获取默认值标签
		defaultTag := structField.Tag.Get("default")
		if defaultTag == "" {
			continue
		}

		// 如果字段为空值，则设置默认值
		if isZeroValue(field) {
			setFieldValue(field, defaultTag)
		}
	}
}

// isZeroValue 检查值是否为零值
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Slice, reflect.Map:
		return v.IsNil()
	default:
		return false
	}
}

// setFieldValue 根据类型设置字段值
func setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var val int64
		fmt.Sscanf(value, "%d", &val)
		field.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var val uint64
		fmt.Sscanf(value, "%d", &val)
		field.SetUint(val)
	case reflect.Float32, reflect.Float64:
		var val float64
		fmt.Sscanf(value, "%f", &val)
		field.SetFloat(val)
	case reflect.Bool:
		var val bool
		fmt.Sscanf(value, "%t", &val)
		field.SetBool(val)
	}
}
