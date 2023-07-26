package array

/*
	第1题：两数之和
		给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
		你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/

// TwoSum 暴力破解
func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

func TwoSum2(nums []int, target int) []int {
	var result []int
	m := make(map[int]int)
	for k, v := range nums {
		if value, exist := m[target-v]; exist {
			result = append(result, value)
			result = append(result, k)
		}
		m[v] = k
	}
	return result // 并返回他们的数组下标。
}
