package array

/*

	题目: 旋转数组

	给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

	说明:
		尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
		要求使用空间复杂度为 O(1) 的 原地 算法。

*/

func Rotate(nums []int, k int) {
	if k >= len(nums) {
		panic("k超出范围")
	}
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

// sort.Reverse()
// 反转
func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]

	}
}
