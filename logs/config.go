package logs

// LogConfig 配置结构
type LogConfig struct {
	TraceOutput   OutputInfo // 追踪日志输出位置
	InfoOutput    OutputInfo // 信息日志输出位置
	WarningOutput OutputInfo // 警告日志输出位置
	ErrorOutput   OutputInfo // 错误日志输出位置
	Flags         int        // 日志标志
}

type OutputInfo struct {
	isConsole    bool
	ConsoleMode  ConsoleMode
	isOutputFile bool
	OutputFile   []string
}

type ConsoleMode string

const (
	DiscardMode ConsoleMode = "discard"
	InfoMode    ConsoleMode = "stdout"
	ErrorMode   ConsoleMode = "stderr"
)
