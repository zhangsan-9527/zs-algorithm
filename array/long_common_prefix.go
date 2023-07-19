package array

import (
	"fmt"
	"strings"
)

/*
	题目: 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
*/

func LongComminPrefix(arr []string) string {
	if len(arr) < 1 {
		return ""
	}

	prefix := arr[0]
	for _, s := range arr {
		for strings.Index(s, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
			fmt.Println(prefix)
		}
	}
	return prefix
}
