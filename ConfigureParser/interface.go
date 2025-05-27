package ConfigureParser

// Interface 配置文件的解析器
type Interface interface {
	ConfigureParser(configPath string, config any)
}
