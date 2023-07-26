package array

import "fmt"

/*
	题目27：移除元素
		给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
		不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
		元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/

func RemoveEle(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

/*
	题目26：删除排序数组中的重复项
	给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次。
	返回移除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/

func OnlyOnce(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return len(nums)
}
