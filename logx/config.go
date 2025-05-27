package logx

import (
	"errors"
	"io"
)

// LogConfig 配置结构
type LogConfig struct {
	TraceOutput   OutputInfo // 追踪日志输出位置
	InfoOutput    OutputInfo // 信息日志输出位置
	DeBugOutput   OutputInfo // 调试日志输出位置
	WarningOutput OutputInfo // 警告日志输出位置
	ErrorOutput   OutputInfo // 错误日志输出位置
	Flags         int        // 日志标志
}

type OutputInfo struct {
	IsConsole    bool
	ConsoleMode  ConsoleMode
	IsOutputFile bool
	OutputFile   []string
	IsCustom     bool       // 是否设置自定义的处理操作
	CustomWriter *io.Writer // 自定义的处理操作
}

type ConsoleMode string

const (
	DiscardMode ConsoleMode = "discard"
	InfoMode    ConsoleMode = "stdout"
	ErrorMode   ConsoleMode = "stderr"
)

type LogPattern int

const (
	Trace LogPattern = iota + 1
	Info
	DeBug
	Warning
	Error
)

// LogConfigAddWriter 设置LogConfig的自定义Writer
func LogConfigAddWriter(config *LogConfig, pattern LogPattern, writer *io.Writer) error {

	var info *OutputInfo

	switch pattern {
	case Trace:
		info = &config.TraceOutput
	case Info:
		info = &config.InfoOutput
	case DeBug:
		info = &config.DeBugOutput
	case Warning:
		info = &config.WarningOutput
	case Error:
		info = &config.ErrorOutput
	default:
		return errors.New("pattern is error")
	}

	if writer != nil {
		info.IsCustom = true
		info.CustomWriter = writer
	} else {
		info.IsCustom = false
		info.CustomWriter = nil
	}

	return nil
}
