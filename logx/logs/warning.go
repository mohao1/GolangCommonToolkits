package logs

import (
	"common-toolkits-v1/logx"
	"fmt"
	"os"
)

// Warning 函数
func Warning(v ...any) {
	logx.WarningLevel.Output(2, fmt.Sprint(v...))
}

func WarningF(info string, v ...any) {
	logx.WarningLevel.Output(2, fmt.Sprintf(info, v...))
}

func Warningln(v any) {
	logx.WarningLevel.Output(2, fmt.Sprintln(v))
}

func WarningFatal(v ...any) {
	logx.WarningLevel.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func WarningFatalf(format string, v ...any) {
	logx.WarningLevel.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func WarningFatalln(v ...any) {
	logx.WarningLevel.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func WarningPanic(v ...any) {
	s := fmt.Sprint(v...)
	logx.WarningLevel.Output(2, s)
	panic(s)
}

func WarningPanicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logx.WarningLevel.Output(2, s)
	panic(s)
}

func WarningPanicln(v ...any) {
	s := fmt.Sprintln(v...)
	logx.WarningLevel.Output(2, s)
	panic(s)
}
