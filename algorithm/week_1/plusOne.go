// 模拟加法进位的思想
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		// 进位，只要有其中一位不是0，就可以完成加1后返回
		if digits[i] != 0 {
			return digits
		}
	}
	// 如果所有位均为0，即类似[9,9,9...], 直接返回一个首位1，其余为0，总长度比原数组大1的数组
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
