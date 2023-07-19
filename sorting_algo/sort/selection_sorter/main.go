package main

import (
	"algo/sorting_algo/sort/base"
	"fmt"
)

func SelectionSorter(data []int) {
	if data == nil || len(data) <= 1 {
		return
	}

	for i := 0; i < len(data); i++ { // 控制选择排序的轮数
		mixIndex := i // 最小元素的索引
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[mixIndex] {
				mixIndex = j
			}
		}
		// 将 data[i] 和最小元素交换
		//tmp := data[i]
		//data[i] = data[mixIndex]
		//data[mixIndex] = tmp
		base.Swap(data, i, mixIndex)
	}

}

func main() {

	data := []int{12, 23, 36, 9, 24, 42}

	SelectionSorter(data)
	fmt.Println(data)

}
