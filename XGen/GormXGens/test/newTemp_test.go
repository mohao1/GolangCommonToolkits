package test

import (
	"common-toolkits-v1/GormXGens/config"
	"common-toolkits-v1/GormXGens/xgen"
	"common-toolkits-v1/logx/logs"
	"fmt"
	"gorm.io/gen"
	"os"
	"testing"
	"text/template"
)

func TestNewTemp(t *testing.T) {

	structInfo := config.TemplateField{
		QueryName:     "UserDo",
		QueryPackage:  "test",
		ModelName:     "User",
		ModelPackage:  "entity",
		ModelLinkPath: "common-toolkits-v1/GormXGens/go_template/entity",
	}

	templateContent, err := os.ReadFile("../template/xgen_template.tpl")
	if err != nil {
		return
	}

	parseTemp, err := template.New("xgen_template").Parse(string(templateContent))
	if err != nil {
		fmt.Printf("解析模板出错: %v\n", err)
		return
	}

	// 执行模板并将结果输出到文件
	file, err := os.Create("./testFile.go")
	if err != nil {
		fmt.Printf("创建文件出错: %v\n", err)
		return
	}
	defer file.Close()

	err = parseTemp.Execute(file, structInfo)
	if err != nil {
		fmt.Printf("执行模板出错: %v\n", err)
		return
	}

}

func TestPath(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		return
	}
	//
	//cmd := exec.Command("go", "list", "-m")
	//output, err := cmd.Output()
	//if err != nil {
	//	return
	//}
	//
	//gomodPath := strings.TrimSpace(string(output))
	//
	fmt.Println("Current Working Directory (via filepath.Abs):", path)
}

func TestReflectXGenModel(t *testing.T) {

	modelObj, err := xgen.ParseStructFromFile("../go_template/entity/entity.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("结构体名称: %s\n", modelObj.ModelName)
	fmt.Println("字段信息:")
	for _, field := range modelObj.FieldMap {
		fmt.Printf("  字段名称: %s, 字段类型: %s, gorm标签: %s, json标签: %s\n", field.FieldName, field.FieldType, field.FieldGormTag, field.FieldJsonTag)
		fmt.Printf("gorm标签的key:%v\n", field.GormTag)
	}

}

func TestXGen(t *testing.T) {
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

func TestLogs(t *testing.T) {
	logs.Error("hello")
}
