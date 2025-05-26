package config

import "gorm.io/gorm"

// DBFunc 扩展函数类型
type DBFunc func(db *gorm.DB)

// TemplateField 模板生成数据类型
type TemplateField struct {
	DoName        string // Do操作名称
	DoPackage     string // Do操作生成位置Package
	ModelName     string // Model结构体的名称
	ModelPackage  string // Model存储路径Package
	ModelLinkPath string //  Model引入路径Path
}

type TemplateFieldExtend struct {
	ParentName   string // 父类名称
	ExtendName   string // 扩展名称
	QueryPackage string // 生成位置Package
}
