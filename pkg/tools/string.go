package tools

import "strings"

// 字符串是否为空
func StringIsEmply(str string) bool {

	if str == "" {
		return true
	}

	return false
}

// 字符串是否全为空白字符
func StringIsWhiteSpace(str string) bool {

	return StringIsEmply(strings.TrimSpace(str))
}

// 字符串是否为nil
func StringIsNil(str *string) bool {
	if str == nil {
		return true
	}

	return false
}

// 字符串是否不为空
func StringIsNotEmpty(str string) bool {
	return !StringIsEmply(str)
}

// 字符串是否为空或全为空白字符
func StringIsEmplyOrWhiteSpace(str string) bool {
	return StringIsEmply(str) || StringIsWhiteSpace(str)
}
