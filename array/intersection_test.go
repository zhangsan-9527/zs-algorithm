package array

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	arr1 := []int{1, 5, 8, 7, 54, 8, 8}
	arr2 := []int{1, 2, 3, 4, 5, 7, 5, 75, 8, 8, 8, 8}
	result := Intersect(arr1, arr2)
	fmt.Println(result)
}

func TestIntersect_Up(t *testing.T) {
	arr1 := []int{1, 5, 8, 7, 54, 8, 8}
	arr2 := []int{1, 2, 3, 4, 5, 7, 5, 75, 8, 8, 8, 8}
	result := IntersectUp(arr1, arr2)
	fmt.Println(result)
}
