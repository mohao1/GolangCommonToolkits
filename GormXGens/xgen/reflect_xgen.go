package xgen

import (
	"common-toolkits-v1/logx/logs"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

// FieldInfo Field字段的数据
type FieldInfo struct {
	FieldName    string            // 字段名称
	FieldType    ast.Expr          // 数据类型
	FieldGormTag reflect.StructTag // Gorm的Tag
	GormTag      GormTag           // Gorm数据信息
	FieldJsonTag reflect.StructTag // Json的Tag
}

// GormTag GORM标签解析结果
type GormTag struct {
	IsPrimaryKey  bool
	IsIndex       bool
	IsUnique      bool
	IndexName     string
	IndexPriority int
}

// ModelObject ModelObject的数据
type ModelObject struct {
	ModelName string                // 结构体的名称
	FieldMap  map[string]*FieldInfo // 结构体的字段
}

func ParseStructFromFile(filePath string) (*ModelObject, error) {
	// 解析文件
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		logs.Error("解析文件时出错:", err)
		return nil, err
	}

	ModelName := ""
	FieldMap := make(map[string]*FieldInfo)

	// 遍历文件中的每个声明
	for _, decl := range file.Decls {
		// 判断是否存在
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		// 遍历每个类型定义
		for _, spec := range genDecl.Specs {

			// 判断结构体是否是存在
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// 获取结构体的类型
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			ModelName = typeSpec.Name.Name
			// 获取字段信息
			for _, field := range structType.Fields.List {
				info := FieldInfo{}
				// 解析字段名称
				if len(field.Names) > 0 {
					info.FieldName = field.Names[0].Name
				}
				// 解析字段类型
				info.FieldType = field.Type
				// 解析jsonTag
				if field.Tag != nil {
					tagValue := field.Tag.Value[1 : len(field.Tag.Value)-1]
					// 解析gorm标签
					info.FieldGormTag = parseTag(tagValue, "gorm")
					// 解析gorm标签配置
					info.GormTag = parseGormTag(string(info.FieldGormTag))
					// 解析json标签
					info.FieldJsonTag = parseTag(tagValue, "json")
				}

				// 存储Field信息
				FieldMap[info.FieldName] = &info
			}
		}
	}
	modelObject := &ModelObject{
		ModelName: ModelName,
		FieldMap:  FieldMap,
	}
	return modelObject, nil
}

// 解析特定类型的标签
func parseTag(tagValue, tagType string) reflect.StructTag {
	tagStart := fmt.Sprintf(`%s:"`, tagType)
	startIdx := strings.Index(tagValue, tagStart)
	if startIdx == -1 {
		return ""
	}
	startIdx += len(tagStart)

	// 查找对应的结束引号
	endIdx := startIdx
	for endIdx < len(tagValue) {
		if tagValue[endIdx] == '"' {
			// 检查是否是转义的引号
			if endIdx > 0 && tagValue[endIdx-1] == '\\' {
				endIdx++
				continue
			}
			break
		}
		endIdx++
	}

	// 确保索引有效
	if endIdx >= len(tagValue) || startIdx > endIdx {
		return ""
	}

	tagStr := tagValue[startIdx:endIdx]

	return reflect.StructTag(tagStr)
}

// ParseGormTag 解析GORM标签，提取主键和索引信息
func parseGormTag(tagValue string) GormTag {
	result := GormTag{}

	if tagValue == "" {
		return result
	}

	// 分割标签中的键值对
	parts := strings.Split(tagValue, ";")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// 检查是否为主键
		if part == "primaryKey" {
			result.IsPrimaryKey = true
			continue
		}

		// 检查是否为索引
		if strings.HasPrefix(part, "index:") || strings.HasPrefix(part, "uniqueIndex:") {
			result.IsIndex = true

			if strings.HasPrefix(part, "uniqueIndex:") {
				result.IsUnique = true
				part = strings.TrimPrefix(part, "uniqueIndex:")
			} else {
				part = strings.TrimPrefix(part, "index:")
			}

			// 解析索引参数
			indexParams := strings.Split(part, ",")
			for _, param := range indexParams {
				param = strings.TrimSpace(param)
				if param == "" {
					continue
				}

				// 格式可能是 name:idx_name 或 priority:N
				if strings.Contains(param, ":") {
					kv := strings.SplitN(param, ":", 2)
					if len(kv) == 2 {
						key, value := kv[0], kv[1]
						switch strings.TrimSpace(key) {
						case "name":
							result.IndexName = strings.TrimSpace(value)
						case "priority":
							if n, err := parseInt(value); err == nil {
								result.IndexPriority = n
							}
						}
					}
				} else if result.IndexName == "" {
					// 如果没有参数名，默认为索引名
					result.IndexName = param
				}
			}
		}
	}

	return result
}

// 辅助函数：将字符串转换为整数
func parseInt(s string) (int, error) {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}
