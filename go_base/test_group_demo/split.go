package test_group_demo

import "strings"

// Split切割字符串
// example:
// abc,b  => [a c]

//func Split(str string, sep string) []string {
//	var ret []string
//	index := strings.Index(str, sep)
//	for index >= 0 {
//		ret = append(ret, str[:index])
//		str = str[index+1:]
//		index = strings.Index(str, sep)
//	}
//	ret = append(ret,str)
//	return ret
//}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return result
}
