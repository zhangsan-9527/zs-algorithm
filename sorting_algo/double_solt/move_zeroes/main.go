package main

import "fmt"

func MoveZeroes(arr []int) {
	newArr := make([]int, len(arr))
	i := 0
	for _, val := range arr {
		if val != 0 {
			newArr[i] = val
			i++
		}
	}

	for k, val := range newArr {
		arr[k] = val
	}

}

// MoveZeroesTow 原地算法: 在原始输入数组上完成的算法没有申请额外的空间
// slow 永远表示下一个不为0的元素的位置
func MoveZeroesTow(arr []int) {
	slow := 0
	for fast, val := range arr {
		if val != 0 { // 减少交换次数
			if fast != slow {
				arr[fast] = arr[slow]
				arr[slow] = val
			}
			slow++
		}
	}

}
func MoveZeroesThrid(arr []int) {
	slow := 0
	for fast := 0; fast < len(arr); fast++ {
		if arr[fast] != 0 { // 减少交换次数
			if fast != slow {
				arr[slow] = arr[fast]
			}
			slow++
		}
	}
	for ; slow < len(arr); slow++ {
		arr[slow] = 0
	}
}
func main() {
	arr := []int{2, 4, 57, 0, 4, 5, 0, 56, 78, 7, 12}
	//MoveZeroes(arr)
	//MoveZeroesTow(arr)
	MoveZeroesThrid(arr)
	fmt.Println(arr)
}
