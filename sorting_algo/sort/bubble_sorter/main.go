package main

import (
	"algo/sorting_algo/sort/base"
	"fmt"
)

func BubbleSorter(data []int) {

	if data == nil || len(data) <= 1 {
		return
	}

	for roind := 1; roind <= len(data); roind++ { // 控制冒泡轮数
		isSwap := false                   // 表示该轮是否进行交换了(减少比较次数)
		compareTimes := len(data) - roind // 每轮比较次数
		for i := 0; i < compareTimes; i++ {
			if data[i] > data[i+1] {
				//tmp := data[i+1]
				//data[i+1] = data[i]
				//data[i] = tmp
				base.Swap(data, i, i+1)
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

func main() {

	data := []int{12, 23, 36, 9, 24, 42}

	BubbleSorter(data)
	fmt.Println(data)

}
