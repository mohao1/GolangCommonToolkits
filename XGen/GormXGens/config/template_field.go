package config

import "gorm.io/gorm"

// DBFunc 扩展函数类型
type DBFunc func(db *gorm.DB) error

// PrimaryKeyData 主键类型
type PrimaryKeyData struct {
	Name   string // 主键名称
	Type   string // 类型
	Column string // Column名称
}

// IndexKeyData 索引类型
type IndexKeyData struct {
	Name   string // 主键名称
	Type   string // 类型
	Column string // Column名称
}

// TemplateField 模板生成数据类型
type TemplateField struct {
	QueryName     string          // Do操作名称
	QueryPackage  string          // Do操作生成位置Package
	ModelName     string          // Model结构体的名称
	ModelPackage  string          // Model存储路径Package
	ModelLinkPath string          //  Model引入路径Path
	PrimaryKey    *PrimaryKeyData // PrimaryKey 主键类型
	IndexList     []IndexKeyData  // IndexKey 索引类型
}

type TemplateFieldExtend struct {
	ParentName   string // 父类名称
	ExtendName   string // 扩展名称
	QueryPackage string // 生成位置Package
}
