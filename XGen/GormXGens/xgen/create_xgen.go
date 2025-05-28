package xgen

import (
	"common-toolkits-v1/ConfigureParser/YamlParser"
	"common-toolkits-v1/XGen/GormXGens/config"
	"common-toolkits-v1/logx/logs"
	"gorm.io/gen"
)

// YamlConfigNewXGen 通过配置文件生成XGen
func YamlConfigNewXGen(yamlConfigPath string) (*XGen, error) {
	yamlParser := YamlParser.NewYamlParser()
	configParser := config.XGenConfigParser{}
	err := yamlParser.ConfigureParser(yamlConfigPath, &configParser)
	if err != nil {
		logs.Errorf("yamlParser err %v", err)
		return nil, err
	}

	xGenConfig := HashYamlConfigToXGenConfig(configParser)

	xGen := NewXGen(xGenConfig)

	return xGen, nil
}

// HashYamlConfigToXGenConfig YamlConfig转换成为XGenConfig
func HashYamlConfigToXGenConfig(configParser config.XGenConfigParser) config.XGenConfig {
	modelPkgPath := configParser.ModelConfig.ModelPkgPath
	dsn := configParser.ModelConfig.DNS
	g := gen.Config{
		ModelPkgPath:      modelPkgPath,
		FieldNullable:     configParser.ModelConfig.FieldNullable,
		FieldCoverable:    configParser.ModelConfig.FieldCoverable,
		FieldSignable:     configParser.ModelConfig.FieldSignable,
		FieldWithIndexTag: configParser.ModelConfig.FieldWithIndexTag,
		FieldWithTypeTag:  configParser.ModelConfig.FieldWithTypeTag,
	}

	xGenConfig := config.XGenConfig{
		GenConfig:    g,
		DNS:          dsn,
		GenModelList: configParser.QueryConfig.GenModelList,
		ProjectPath:  configParser.QueryConfig.ProjectPath,
		QueryPkgPath: configParser.QueryConfig.QueryPkgPath,
		ModelPkgPath: configParser.QueryConfig.ModelPkgPath,
	}

	return xGenConfig
}

// CreateXGenYamlConfig 快速创建XGen生成代码的工具 - 输入配置文件即可
func CreateXGenYamlConfig(yamlConfigPath string) {
	xGen, err := YamlConfigNewXGen(yamlConfigPath)
	if err != nil {
		logs.Errorf("CreateXGenYamlConfig err %v", err)
		return
	}
	err = xGen.CreateXGen()
	if err != nil {
		logs.Errorf("CreateXGen err %v", err)
		return
	}
}

// CreateXGenYamlConfigCustomTemplate 快速创建XGen生成自定义的模板代码的工具 - 输入配置文件、函数、模板路径即可
func CreateXGenYamlConfigCustomTemplate(yamlConfigPath, parentTemplatePath, extendTemplatePath string, f CreateXGenCustom) {
	xGen, err := YamlConfigNewXGen(yamlConfigPath)
	if err != nil {
		logs.Errorf("CreateXGenYamlConfig err %v", err)
		return
	}
	err = xGen.CreateXGenCustomTemplate(parentTemplatePath, extendTemplatePath, f)
	if err != nil {
		logs.Errorf("CreateXGen err %v", err)
		return
	}
}
