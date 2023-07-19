package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {

	nums := []int{12, 45, 7, 45, 78, 42, 54, 78}
	Rotate(nums, 2)
	fmt.Println(nums)

}
