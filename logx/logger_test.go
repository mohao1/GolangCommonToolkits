package logx

import (
	"common-toolkits-v1/ConfigureParser/YamlParser"
	"fmt"
	"log"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	TraceLevel.Println("Trace-输出")
	InfoLevel.Println("Info-输出")
	WarningLevel.Println("Warning-输出")
	ErrorLevel.Println("Error-输出")
}

func TestInitLogger(t *testing.T) {

	err := InitLogger(&LogConfig{
		TraceOutput: OutputInfo{
			IsConsole:   true,
			ConsoleMode: InfoMode,
		},
		InfoOutput: OutputInfo{
			IsConsole:   true,
			ConsoleMode: InfoMode,
		},
		WarningOutput: OutputInfo{
			IsConsole:    true,
			ConsoleMode:  InfoMode,
			IsOutputFile: true,
			OutputFile: []string{
				"./warning.log",
			},
		},
		ErrorOutput: OutputInfo{
			IsConsole:    true,
			ConsoleMode:  ErrorMode,
			IsOutputFile: true,
			OutputFile: []string{
				"./error.log",
			},
		},
		Flags: log.Ldate | log.Ltime | log.Lshortfile,
	})
	if err != nil {
		return
	}

	TraceLevel.Println("Trace-输出")
	InfoLevel.Println("Info-输出")
	WarningLevel.Println("Warning-输出")
	ErrorLevel.Println("Error-输出")
}

func TestConfigYaml(t *testing.T) {
	yamlParser := YamlParser.NewYamlParser()
	logConfigParser := LogConfigParser{}
	err := yamlParser.ConfigureParser("./logger_config.yaml", &logConfigParser)
	if err != nil {
		return
	}
	fmt.Println(logConfigParser)
}
