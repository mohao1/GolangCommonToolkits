package script_test

import (
	"common-toolkits-v1/GormXGens/config"
	"common-toolkits-v1/GormXGens/xgen"
	"gorm.io/gen"
	"testing"
)

func TestXgen(t *testing.T) {
	modelPath := "model/entity"
	dsn := "root:123456@tcp(localhost:3306)/health_system?charset=utf8mb4&parseTime=True&loc=Local"
	g := gen.Config{
		ModelPkgPath:      modelPath,
		Mode:              gen.WithQueryInterface | gen.WithDefaultQuery,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	}

	xGenConfig := config.XGenConfig{
		GenConfig:    g,
		DNS:          dsn,
		GenModelList: []string{"users", "apply_link_doctor"},
		ProjectPath:  "golang-common-toolkits",
		QueryPkgPath: "../model/query",
		ModelPkgPath: "../model/entity",
	}

	Xgen := xgen.NewXGen(xGenConfig)
	err := Xgen.CreateXGenModel()
	if err != nil {
		return
	}

	err = Xgen.CreateXGenQuery()
	if err != nil {
		return
	}
}

func TestXgenCreateXGen(t *testing.T) {
	modelPath := "model/entity"
	dsn := "root:123456@tcp(localhost:3306)/health_system?charset=utf8mb4&parseTime=True&loc=Local"
	g := gen.Config{
		ModelPkgPath:      modelPath,
		Mode:              gen.WithQueryInterface | gen.WithDefaultQuery,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	}

	xGenConfig := config.XGenConfig{
		GenConfig:    g,
		DNS:          dsn,
		GenModelList: []string{"users", "apply_link_doctor"},
		ProjectPath:  "golang-common-toolkits",
		QueryPkgPath: "../model/query",
		ModelPkgPath: "../model/entity",
	}

	Xgen := xgen.NewXGen(xGenConfig)
	err := Xgen.CreateXGen()
	if err != nil {
		return
	}
}
