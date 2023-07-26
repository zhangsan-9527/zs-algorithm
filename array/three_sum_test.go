package array

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-4, -3, -1, 0, 1, 2, 3, 4}
	result := ThreeSum(nums)
	fmt.Println(result)
}
