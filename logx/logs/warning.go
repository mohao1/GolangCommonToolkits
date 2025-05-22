package logs

import (
	"common-toolkits-v1/logx"
)

// Warning 函数
func Warning(v ...any) {
	logx.WarningLevel.Print(v)
}

func WarningF(info string, v ...any) {
	logx.WarningLevel.Printf(info, v)
}

func Warningln(v any) {
	logx.WarningLevel.Println(v)
}

func WarningFatal(v ...any) {
	logx.WarningLevel.Fatal(v)
}

func WarningFatalf(format string, v ...any) {
	logx.WarningLevel.Fatalf(format)
}

func WarningFatalln(v ...any) {
	logx.WarningLevel.Fatalln(v)
}

func WarningPanic(v ...any) {
	logx.WarningLevel.Panic(v)
}

func WarningPanicf(format string, v ...any) {
	logx.WarningLevel.Panicf(format, v)
}

func WarningPanicln(v ...any) {
	logx.WarningLevel.Panicln(v)
}
