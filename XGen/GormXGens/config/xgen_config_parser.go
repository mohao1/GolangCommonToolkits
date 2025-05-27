package config

// XGenConfigParser Gen配置解析
type XGenConfigParser struct {
	ModelConfig ModelConfig `yaml:"ModelConfig" required:"true"`
	QueryConfig QueryConfig `yaml:"QueryConfig" required:"true"`
}

// ModelConfig Model生成配置
type ModelConfig struct {
	DNS               string `yaml:"DNS" required:"true"`
	ModelPkgPath      string `yaml:"ModelPkgPath" required:"true"`
	FieldNullable     bool   `yaml:"FieldNullable" default:"true"`
	FieldCoverable    bool   `yaml:"FieldCoverable" default:"false"`
	FieldSignable     bool   `yaml:"FieldSignable " default:"true"`
	FieldWithIndexTag bool   `yaml:"FieldWithIndexTag" default:"true"`
	FieldWithTypeTag  bool   `yaml:"FieldWithTypeTag" default:"true"`
}

// QueryConfig Query生成配置
type QueryConfig struct {
	GenModelList []string `yaml:"GenModelList" required:"true"`
	ProjectPath  string   `yaml:"ProjectPath" required:"true"`
	QueryPkgPath string   `yaml:"QueryPkgPath" required:"true"`
	ModelPkgPath string   `yaml:"ModelPkgPath" required:"true"`
}
