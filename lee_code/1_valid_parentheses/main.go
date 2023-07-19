package main

import (
	"fmt"
	"strings"
)

func IsMatch(c1, c2 string) bool {
	if c1 == "(" {
		return c2 == ")"
	} else if c1 == "[" {
		return c2 == "]"
	} else if c1 == "{" {
		return c2 == "}"
	}
	return false
}

func Valid_Parentheses(s string) bool {
	sb := strings.Split(s, "")
	if len(sb)%2 != 0 {
		return false
	}
	for {
		// 匹配完成
		if len(sb) == 0 {
			return true
		} else {
			count := len(sb) / 2
			c1 := sb[count-1]
			c2 := sb[count]
			result := IsMatch(c1, c2)
			if result {
				sb = append(sb[:count-1], sb[count+1:]...)
				fmt.Println(sb)
			} else {
				return false
			}
		}
	}

}

func main() {
	fmt.Println(Valid_Parentheses("{[{}]}"))
}
