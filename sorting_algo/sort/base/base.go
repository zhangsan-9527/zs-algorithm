package base

func Swap(nums []int, i, j int) {

	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp

}
