package array

import (
	"sort"
)

/*
	第15题：三数之和
		给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。注意：答案中不可以包含重复的三元组。

*/

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		l := i + 1
		r := len(nums) - 1

		if nums[i] > 0 {
			break
		}

		if i == 0 || nums[i] != nums[i-1] {
			for l < r {
				if nums[l]+nums[r] == target {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return res
}
