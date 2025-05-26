package config

import (
	"gorm.io/gen"
)

// XGenConfig XGen配置
type XGenConfig struct {
	GenConfig    gen.Config // Gen配置
	DNS          string     // 数据库的地址
	GenModelList []string   // 需要生成表的列表
	ProjectPath  string     // 项目文件夹的名称
	QueryPkgPath string     // 生成的Query操作的文件路径
	ModelPkgPath string     // Model的生成路径
}
