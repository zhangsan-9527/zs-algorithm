package main

import (
	"algo/sorting_algo/sort/base"
	"fmt"
)

func InsertionSort(data []int) {
	if data == nil || len(data) <= 1 {
		return
	}

	for i := 1; i < len(data); i++ { // 控制插入排序的轮数
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				base.Swap(data, j, j-1)
			} else {
				break
			}
		}
	}
}

func InsertionSort2(data []int) {
	if data == nil || len(data) <= 1 {
		return
	}

	for i := 1; i < len(data); i++ { // 控制插入排序的轮数
		tmp := data[i]
		var j int
		for j = i; j > 0; j-- {
			if tmp < data[j-1] {
				// 将较大的元素总是向右移动
				data[j] = data[j-1]
			} else {
				break
			}
		}
		// 找到 i 对应的元素需要插入的位置
		data[j] = tmp
	}
}

func main() {

	data := []int{12, 23, 36, 9, 24, 42}

	InsertionSort2(data)
	fmt.Println(data)

}
