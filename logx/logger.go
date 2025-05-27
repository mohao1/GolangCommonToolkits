package logx

import (
	"common-toolkits-v1/ConfigureParser/YamlParser"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var defaultConfig = LogConfig{
	TraceOutput: OutputInfo{
		IsConsole:   true,
		ConsoleMode: DiscardMode,
	},
	InfoOutput: OutputInfo{
		IsConsole:   true,
		ConsoleMode: InfoMode,
	},
	DeBugOutput: OutputInfo{
		IsConsole:   true,
		ConsoleMode: InfoMode,
	},
	WarningOutput: OutputInfo{
		IsConsole:   true,
		ConsoleMode: InfoMode,
	},
	ErrorOutput: OutputInfo{
		IsConsole:   true,
		ConsoleMode: ErrorMode,
	},
	Flags: log.Ldate | log.Ltime | log.Llongfile,
}

var (
	TraceLevel   *log.Logger
	InfoLevel    *log.Logger
	WarningLevel *log.Logger
	ErrorLevel   *log.Logger
	DeBugLevel   *log.Logger

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
		TraceLevel = log.New(io.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Llongfile)
		InfoLevel = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		DeBugLevel = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
		WarningLevel = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
		ErrorLevel = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	}
}

// InitLoggerYaml 解析Yaml配置文件配置Config
func InitLoggerYaml(configYamlPath string) error {
	yamlParser := YamlParser.NewYamlParser()
	logConfigParser := LogConfigParser{}
	err := yamlParser.ConfigureParser("./logger_config.yaml", &logConfigParser)
	if err != nil {
		fmt.Println("Failed to parse logger config")
		return err
	}
	// 解析数据 文件解析数据 => 配置文件数据
	logConfig := &LogConfig{
		TraceOutput:   HashOutput(logConfigParser.TraceOutput),
		InfoOutput:    HashOutput(logConfigParser.InfoOutput),
		DeBugOutput:   HashOutput(logConfigParser.DeBugOutput),
		WarningOutput: HashOutput(logConfigParser.WarningOutput),
		ErrorOutput:   HashOutput(logConfigParser.ErrorOutput),
		Flags:         HashFlags(logConfigParser.Flags),
	}

	// 初始化数据的信息
	err = InitLogger(logConfig)
	if err != nil {
		fmt.Println("Failed to initialize logger config")
		return err
	}

	return nil
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

	if config.Flags == 0 {
		e := fmt.Sprint("config.Flags is 0")
		return errors.New(e)
	}

	// Trace - 追踪日志
	TraceLevel, err = createLogger(config.TraceOutput, "TRACE: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize trace logger: %v", err)
	}

	// Info - 信息日志
	InfoLevel, err = createLogger(config.InfoOutput, "INFO: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize info logger: %v", err)
	}

	// Warning - 调试日志
	DeBugLevel, err = createLogger(config.DeBugOutput, "DEBUG: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize warning logger: %v", err)
	}

	// Warning - 警告日志
	WarningLevel, err = createLogger(config.WarningOutput, "WARNING: ", config.Flags)
	if err != nil {
		return fmt.Errorf("failed to initialize warning logger: %v", err)
	}

	// Error - 错误日志
	ErrorLevel, err = createLogger(config.ErrorOutput, "ERROR: ", config.Flags)
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

	if outputInfo.IsConsole && isMode {
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

	if outputInfo.IsOutputFile {
		for _, output := range outputInfo.OutputFile {
			writer, err := createWriter(ConsoleMode(output))
			if err != nil {
				return nil, err
			}
			writerList = append(writerList, writer)
		}
	}

	// 自定义的log的操作
	if outputInfo.IsCustom && outputInfo.CustomWriter != nil {
		writerList = append(writerList, *outputInfo.CustomWriter)
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
func reset() {
	mu.Lock()
	defer mu.Unlock()

	TraceLevel = nil
	InfoLevel = nil
	WarningLevel = nil
	ErrorLevel = nil
	initialized = false
}
