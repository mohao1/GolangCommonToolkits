package logs

import "common-toolkits-v1/logx"

// DeBug 函数
func DeBug(v ...any) {
	logx.DeBugLevel.Print(v)
}

func DeBugf(info string, v ...any) {
	logx.DeBugLevel.Printf(info, v)
}

func DeBugln(v any) {
	logx.DeBugLevel.Println(v)
}

func DeBugFatal(v ...any) {
	logx.DeBugLevel.Fatal(v)
}

func DeBugFatalf(format string, v ...any) {
	logx.DeBugLevel.Fatalf(format)
}

func DeBugFatalln(v ...any) {
	logx.DeBugLevel.Fatalln(v)
}

func DeBugPanic(v ...any) {
	logx.DeBugLevel.Panic(v)
}

func DeBugPanicf(format string, v ...any) {
	logx.DeBugLevel.Panicf(format, v)
}

func DeBugPanicln(v ...any) {
	logx.DeBugLevel.Panicln(v)
}
