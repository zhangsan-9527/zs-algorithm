package array

import "sort"

/*
	题目: 给定两个数组，编写一个函数来计算它们的交集。

	进阶: 如果给定的数组已经排好序呢？将如何优化你的算法呢？
*/

/*
	首先拿到这道题，我们基本马上可以想到，此题可以看成是一道传统的映射题（map映射），
	为什么可以这样看呢，因为我们需找出两个数组的交集元素，同时应与两个数组中出现的次数一致。
	这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>。剩下的就是顺利成章的解题。
*/

func Intersect(arr1 []int, arr2 []int) []int {
	mp := map[int]int{}
	for _, v := range arr1 {
		// 遍历arr1, 初始化map
		mp[v] += 1
	}

	k := 0
	for _, v := range arr2 {
		if mp[v] > 0 {
			mp[v] -= 1
			arr2[k] = v
			k++
		}
	}
	return arr2[:k]

}

// IntersectUp 进阶: 如果给定的数组已经排好序呢？将如何优化你的算法呢？
func IntersectUp(arr1, arr2 []int) []int {
	i, j, k := 0, 0, 0
	// 顺序排序
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			arr1[k] = arr1[i]
			i++
			j++
			k++
		}
	}
	return arr1[:k]
}
