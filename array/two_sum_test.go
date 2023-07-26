package array

import (
	"fmt"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	nums := []int{1, 8, 2, 7, 3, 6}
	tag := 9
	result := TwoSum2(nums, tag)
	fmt.Println(result)

}
