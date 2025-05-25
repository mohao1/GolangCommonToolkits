package logs

import (
	"common-toolkits-v1/logx"
	"fmt"
	"os"
)

// Info 函数
func Info(v ...any) {
	logx.InfoLevel.Output(2, fmt.Sprint(v...))
}

func Infof(info string, v ...any) {
	logx.InfoLevel.Output(2, fmt.Sprintf(info, v...))
}

func Infoln(v any) {
	logx.InfoLevel.Output(2, fmt.Sprintln(v))
}

func InfoFatal(v ...any) {
	logx.InfoLevel.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func InfoFatalf(format string, v ...any) {
	logx.InfoLevel.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func InfoFatalln(v ...any) {
	logx.InfoLevel.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func InfoPanic(v ...any) {
	s := fmt.Sprint(v...)
	logx.InfoLevel.Output(2, s)
	panic(s)
}

func InfoPanicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logx.InfoLevel.Output(2, s)
	panic(s)
}

func InfoPanicln(v ...any) {
	s := fmt.Sprintln(v...)
	logx.InfoLevel.Output(2, s)
	panic(s)
}
