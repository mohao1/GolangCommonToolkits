package logs

import "common-toolkits-v1/logx"

// Info 函数
func Info(v ...any) {
	logx.InfoLevel.Print(v)
}

func Infof(info string, v ...any) {
	logx.InfoLevel.Printf(info, v)
}

func Infoln(v any) {
	logx.InfoLevel.Println(v)
}

func InfoFatal(v ...any) {
	logx.InfoLevel.Fatal(v)
}

func InfoFatalf(format string, v ...any) {
	logx.InfoLevel.Fatalf(format)
}

func InfoFatalln(v ...any) {
	logx.InfoLevel.Fatalln(v)
}

func InfoPanic(v ...any) {
	logx.InfoLevel.Panic(v)
}

func InfoPanicf(format string, v ...any) {
	logx.InfoLevel.Panicf(format, v)
}

func InfoPanicln(v ...any) {
	logx.InfoLevel.Panicln(v)
}
