package logs

import "common-toolkits-v1/logx"

// Trace 函数
func Trace(v ...any) {
	logx.TraceLevel.Print(v)
}

func Tracef(info string, v ...any) {
	logx.TraceLevel.Printf(info, v)
}

func Traceln(v any) {
	logx.TraceLevel.Println(v)
}

func TraceFatal(v ...any) {
	logx.TraceLevel.Fatal(v)
}

func TraceFatalf(format string, v ...any) {
	logx.TraceLevel.Fatalf(format)
}

func TraceFatalln(v ...any) {
	logx.TraceLevel.Fatalln(v)
}

func TracePanic(v ...any) {
	logx.TraceLevel.Panic(v)
}

func TracePanicf(format string, v ...any) {
	logx.TraceLevel.Panicf(format, v)
}

func TracePanicln(v ...any) {
	logx.TraceLevel.Panicln(v)
}
