package main

import (
	"fmt"
	"strings"
)

func ReveString(s string) string {
	//strings.Fields(s) 结果是[hellooo]
	strArr := strings.Split(s, "") // 结果是[h e l l o o o]
	fmt.Println(strArr)
	newArr := make([]string, len(strArr))

	j := len(strArr) - 1
	for i := 0; i < len(strArr); i++ {
		newArr[j] = strArr[i]
		j--
	}

	return strings.Join(newArr, "")
}

func ReveString2(s string) string {
	//strings.Fields(s) 结果是[hellooo]
	strArr := strings.Split(s, "") // 结果是[h e l l o o o]
	right := len(strArr) - 1
	left := 0
	tmp := ""
	for {
		if left < right {
			tmp = strArr[right]
			strArr[right] = strArr[left]
			strArr[left] = tmp
			left++
			right--
		} else {
			return strings.Join(strArr, "")
		}

	}

}

func main() {
	s := "hellooo"
	fmt.Println(ReveString2(s))
}
