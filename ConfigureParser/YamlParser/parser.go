package YamlParser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

type YamlParser struct {
}

func NewYamlParser() *YamlParser {
	return &YamlParser{}
}

// ConfigureParser 解析配置文件
func (y *YamlParser) ConfigureParser(configPath string, config any) error {

	// 判断是否是指针的数据
	if !isPointer(config) {
		return fmt.Errorf("config must be a pointer")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}

	// 默认值的操作
	setDefaults(reflect.ValueOf(config).Elem())

	err = validateRequiredFields(config)
	if err != nil {
		return err
	}

	return nil
}

// 判断接口值是否为指针类型
func isPointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}
