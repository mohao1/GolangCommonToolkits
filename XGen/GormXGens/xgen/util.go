package xgen

import "unicode"

// lowercaseFirst 首字母小写的操作
func lowercaseFirst(s string) string {
	if s == "" {
		return s
	}
	// 将字符串转为 rune 切片以便处理多字节字符
	runes := []rune(s)
	// 把第一个字符转换为小写
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
