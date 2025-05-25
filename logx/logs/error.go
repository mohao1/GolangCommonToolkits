package logs

import (
	"common-toolkits-v1/logx"
	"fmt"
	"os"
)

// Error 函数
func Error(v ...any) {
	logx.ErrorLevel.Output(2, fmt.Sprint(v...))
}

func Errorf(info string, v ...any) {
	logx.ErrorLevel.Output(2, fmt.Sprintf(info, v...))
}

func Errorln(v any) {
	logx.ErrorLevel.Output(2, fmt.Sprintln(v))
}

func ErrorFatal(v ...any) {
	logx.ErrorLevel.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func ErrorFatalf(format string, v ...any) {
	logx.ErrorLevel.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func ErrorFatalln(v ...any) {
	logx.ErrorLevel.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func ErrorPanic(v ...any) {
	s := fmt.Sprint(v...)
	logx.ErrorLevel.Output(2, s)
	panic(s)
}

func ErrorPanicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logx.ErrorLevel.Output(2, s)
	panic(s)
}

func ErrorPanicln(v ...any) {
	s := fmt.Sprintln(v...)
	logx.ErrorLevel.Output(2, s)
	panic(s)
}
