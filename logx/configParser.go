package logx

import "log"

type OutputConfigParser struct {
	IsConsole    bool        `yaml:"IsConsole" default:"false"`
	ConsoleMode  ConsoleMode `yaml:"ConsoleMode" default:"discard"`
	IsOutputFile bool        `yaml:"IsOutputFile" default:"false"`
	OutputFile   []string    `yaml:"OutputFile" default:"[]"`
}

// LogConfigParser 解析配置文件结构
type LogConfigParser struct {
	ModeType      Environment        `yaml:"ModeType" default:"prod"`
	TraceOutput   OutputConfigParser `yaml:"TraceOutput"`
	InfoOutput    OutputConfigParser `yaml:"InfoOutput"`
	DeBugOutput   OutputConfigParser `yaml:"DeBugOutput"`
	WarningOutput OutputConfigParser `yaml:"WarningOutput"`
	ErrorOutput   OutputConfigParser `yaml:"ErrorOutput"`
	Flags         []Flag             `yaml:"Flags" default:"[]"`
}

// Flag 输出文本格式类型
type Flag string

const (
	Ldate         Flag = "Ldate"
	Ltime         Flag = "Ltime"
	Lmicroseconds Flag = "Lmicroseconds"
	Llongfile     Flag = "Llongfile"
	Lshortfile    Flag = "Lshortfile"
	LUTC          Flag = "LUTC"
	Lmsgprefix    Flag = "Lmsgprefix"
	LstdFlags     Flag = "LstdFlags"
)

// Environment 模式类型
type Environment string

const (
	Development Environment = "dev"
	Testing     Environment = "test"
	Production  Environment = "prod"
	Debug       Environment = "debug"
	Sandbox     Environment = "sandbox"
	Staging     Environment = "staging"
)

// HashFlags 解析Flags信息
func HashFlags(flags []Flag) int {
	newFlags := 0
	for k, flag := range flags {
		data := 0
		switch flag {
		case Ldate:
			data = log.Ldate
		case Ltime:
			data = log.Ltime
		case Lmicroseconds:
			data = log.Lmicroseconds
		case Llongfile:
			data = log.Llongfile
		case Lshortfile:
			data = log.Lshortfile
		case LUTC:
			data = log.LUTC
		case Lmsgprefix:
			data = log.Lmsgprefix
		case LstdFlags:
			data = log.LstdFlags
		}
		if k == 0 {
			newFlags = data
		} else {
			newFlags = newFlags | data
		}
	}
	return newFlags
}

// HashOutput 解析OutPut配置文件
func HashOutput(outPut OutputConfigParser) OutputInfo {
	return OutputInfo{
		IsConsole:    outPut.IsConsole,
		ConsoleMode:  outPut.ConsoleMode,
		IsOutputFile: outPut.IsOutputFile,
		OutputFile:   outPut.OutputFile,
	}
}
