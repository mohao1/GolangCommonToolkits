package YamlParser

import (
	"fmt"
	"reflect"
	"strings"
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

func validateRequiredFields(data interface{}) error {
	value := reflect.ValueOf(data)

	// 处理指针
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil // 允许nil指针
		}
		value = value.Elem()
	}

	// 只处理结构体
	if value.Kind() != reflect.Struct {
		return nil
	}
	typ := value.Type()

	// 遍历结构体的每个字段
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)

		// 提取YAML标签和自定义标签
		yamlTag := field.Tag.Get("yaml")
		requiredTag := field.Tag.Get("required")

		// 解析YAML标签获取字段名
		fieldName := yamlTag
		if idx := strings.Index(yamlTag, ","); idx != -1 {
			fieldName = yamlTag[:idx]
		}

		// 检查是否为必需字段
		if requiredTag == "true" || requiredTag == "1" {
			// 如果是必需字段且为空，则报错
			if isZeroValue(fieldValue) {
				return fmt.Errorf("缺少必需字段: %s", fieldName)
			}
		}

		// 递归处理嵌套结构体
		if fieldValue.Kind() == reflect.Struct {
			if err := validateRequiredFields(fieldValue.Addr().Interface()); err != nil {
				return err
			}
		} else if fieldValue.Kind() == reflect.Ptr &&
			!fieldValue.IsNil() &&
			fieldValue.Type().Elem().Kind() == reflect.Struct {
			if err := validateRequiredFields(fieldValue.Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}
