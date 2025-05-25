package logs

import (
	"common-toolkits-v1/logx"
	"fmt"
	"os"
)

// DeBug 函数
func DeBug(v ...any) {
	logx.DeBugLevel.Output(2, fmt.Sprint(v...))
}

func DeBugf(info string, v ...any) {
	logx.DeBugLevel.Output(2, fmt.Sprintf(info, v...))
}

func DeBugln(v any) {
	logx.DeBugLevel.Output(2, fmt.Sprintln(v))
}

func DeBugFatal(v ...any) {
	logx.DeBugLevel.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func DeBugFatalf(format string, v ...any) {
	logx.DeBugLevel.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func DeBugFatalln(v ...any) {
	logx.DeBugLevel.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func DeBugPanic(v ...any) {
	s := fmt.Sprint(v...)
	logx.DeBugLevel.Output(2, s)
	panic(s)
}

func DeBugPanicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logx.DeBugLevel.Output(2, s)
	panic(s)
}

func DeBugPanicln(v ...any) {
	s := fmt.Sprintln(v...)
	logx.DeBugLevel.Output(2, s)
	panic(s)
}
