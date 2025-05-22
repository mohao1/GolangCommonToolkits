package logs

import "common-toolkits-v1/logx"

// Error 函数
func Error(v ...any) {
	logx.ErrorLevel.Print(v)
}

func Errorf(info string, v ...any) {
	logx.ErrorLevel.Printf(info, v)
}

func Errorln(v any) {
	logx.ErrorLevel.Println(v)
}

func ErrorFatal(v ...any) {
	logx.ErrorLevel.Fatal(v)
}

func ErrorFatalf(format string, v ...any) {
	logx.ErrorLevel.Fatalf(format)
}

func ErrorFatalln(v ...any) {
	logx.ErrorLevel.Fatalln(v)
}

func ErrorPanic(v ...any) {
	logx.ErrorLevel.Panic(v)
}

func ErrorPanicf(format string, v ...any) {
	logx.ErrorLevel.Panicf(format, v)
}

func ErrorPanicln(v ...any) {
	logx.ErrorLevel.Panicln(v)
}
