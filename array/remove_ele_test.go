package array

import (
	"fmt"
	"testing"
)

func TestRemoveEle(t *testing.T) {
	nums := []int{1, 1, 5, 4, 4, 5, 4, 34}
	val := 5
	result := RemoveEle(nums, val)
	fmt.Println(result)
}

func TestOnlyOnce(t *testing.T) {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	result := OnlyOnce(nums)
	fmt.Println(result)
}
