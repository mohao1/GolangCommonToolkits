package logs

import (
	"common-toolkits-v1/logx"
	"fmt"
	"os"
)

// Trace 函数
func Trace(v ...any) {
	logx.TraceLevel.Output(2, fmt.Sprint(v...))
}

func Tracef(info string, v ...any) {
	logx.TraceLevel.Output(2, fmt.Sprintf(info, v...))
}

func Traceln(v any) {
	logx.TraceLevel.Output(2, fmt.Sprintln(v))
}

func TraceFatal(v ...any) {
	logx.TraceLevel.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func TraceFatalf(format string, v ...any) {
	logx.TraceLevel.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func TraceFatalln(v ...any) {
	logx.TraceLevel.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func TracePanic(v ...any) {
	s := fmt.Sprint(v...)
	logx.TraceLevel.Output(2, s)
	panic(s)
}

func TracePanicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logx.TraceLevel.Output(2, s)
	panic(s)
}

func TracePanicln(v ...any) {
	s := fmt.Sprintln(v...)
	logx.TraceLevel.Output(2, s)
	panic(s)
}
