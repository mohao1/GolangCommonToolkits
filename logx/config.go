package logx

import "io"

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
	isConsole    bool
	ConsoleMode  ConsoleMode
	isOutputFile bool
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
