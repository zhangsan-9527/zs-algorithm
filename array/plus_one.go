package array

/*
	第66题：加一
	给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
	最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。

*/

func PlusOne(digits []int) []int {
	var result []int
	addon := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += addon
		addon = 0
		if i == len(digits)-1 {
			digits[i]++
		}
		if digits[i] == 10 {
			addon = 1
			digits[i] = digits[i] % 10
		}
	}
	if addon == 1 {
		result = make([]int, 1)
		result[0] = 1
		result = append(result, digits...)
	} else {
		result = digits
	}

	return result

}
