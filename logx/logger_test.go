package logx

import (
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
			isConsole:   true,
			ConsoleMode: InfoMode,
		},
		InfoOutput: OutputInfo{
			isConsole:   true,
			ConsoleMode: InfoMode,
		},
		WarningOutput: OutputInfo{
			isConsole:    true,
			ConsoleMode:  InfoMode,
			isOutputFile: true,
			OutputFile: []string{
				"./warning.log",
			},
		},
		ErrorOutput: OutputInfo{
			isConsole:    true,
			ConsoleMode:  ErrorMode,
			isOutputFile: true,
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
