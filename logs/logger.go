package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var defaultConfig = LogConfig{
	TraceOutput: OutputInfo{
		isConsole:   true,
		ConsoleMode: DiscardMode,
	},
	InfoOutput: OutputInfo{
		isConsole:   true,
		ConsoleMode: InfoMode,
	},
	WarningOutput: OutputInfo{
		isConsole:   true,
		ConsoleMode: InfoMode,
	},
	ErrorOutput: OutputInfo{
		isConsole:   true,
		ConsoleMode: ErrorMode,
	},
	Flags: log.Ldate | log.Ltime | log.Lshortfile,
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger

	initialized bool
	mu          sync.Mutex
)

/*
*	io.Discard  忽略
*	os.Stdout  正常输出
*	os.Stderr 错误输出
 */
func init() {
	if err := InitLogger(nil); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialize default logger: %v\n", err)
		// 回退到基本日志配置
		Trace = log.New(io.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

// InitLogger 初始化日志系统（可在运行时重新配置）
func InitLogger(config *LogConfig) error {
	mu.Lock()
	defer mu.Unlock()

	// 使用默认配置
	if config == nil {
		config = &defaultConfig
	}

	// 初始化日志输出
	var err error

	// Trace - 追踪日志
	Trace, err = createLogger(config.TraceOutput, "TRACE: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize trace logger: %v", err)
	}

	// Info - 信息日志
	Info, err = createLogger(config.InfoOutput, "INFO: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize info logger: %v", err)
	}

	// Warning - 警告日志
	Warning, err = createLogger(config.WarningOutput, "WARNING: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize warning logger: %v", err)
	}

	// Error - 错误日志
	Error, err = createLogger(config.ErrorOutput, "ERROR: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize error logger: %v", err)
	}

	initialized = true
	return nil
}

// createLogger 创建日志记录器
func createLogger(outputInfo OutputInfo, prefix string, flags int) (*log.Logger, error) {

	var writerList []io.Writer

	// 类型校验
	isMode := outputInfo.ConsoleMode == DiscardMode ||
		outputInfo.ConsoleMode == InfoMode ||
		outputInfo.ConsoleMode == ErrorMode

	if outputInfo.isConsole && isMode {
		writer, err := createWriter(outputInfo.ConsoleMode)
		if err != nil {
			return nil, err
		}
		writerList = append(writerList, writer)
	} else {
		writer, err := createWriter(DiscardMode)
		if err != nil {
			return nil, err
		}
		writerList = append(writerList, writer)
	}

	if outputInfo.isOutputFile {
		for _, output := range outputInfo.OutputFile {
			writer, err := createWriter(ConsoleMode(output))
			if err != nil {
				return nil, err
			}
			writerList = append(writerList, writer)
		}
	}

	return log.New(io.MultiWriter(writerList...), prefix, flags), nil
}

func createWriter(outputPath ConsoleMode) (io.Writer, error) {
	switch outputPath {
	case DiscardMode:
		return io.Discard, nil
	case InfoMode:
		return os.Stdout, nil
	case ErrorMode:
		return os.Stderr, nil
	default:
		// 确保目录存在
		filePath := string(outputPath)
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %v", err)
		}

		// 打开或创建文件
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %v", err)
		}

		return file, nil
	}
}

// Reset 重置日志配置
func Reset() {
	mu.Lock()
	defer mu.Unlock()

	Trace = nil
	Info = nil
	Warning = nil
	Error = nil
	initialized = false
}
