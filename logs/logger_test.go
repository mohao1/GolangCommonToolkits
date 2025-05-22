package logs

import (
	"log"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	Trace.Println("Trace-输出")
	Info.Println("Info-输出")
	Warning.Println("Warning-输出")
	Error.Println("Error-输出")
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

	Trace.Println("Trace-输出")
	Info.Println("Info-输出")
	Warning.Println("Warning-输出")
	Error.Println("Error-输出")
}
