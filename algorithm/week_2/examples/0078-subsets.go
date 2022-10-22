// 回溯, 类似题目可以画一棵选择状态树进行描述
func subsets(nums []int) [][]int {
	ans := [][]int{}
	chosen := []int{}

	var recur func(int)
	recur = func(i int) {
		// 临界条件，下标超出，当前方案已经全部作出选择，将chosen加入ans，注意copy一份chosen
		if i == len(nums) {
			tmp := make([]int, len(chosen))
			copy(tmp, chosen)
			ans = append(ans, tmp)
			return
		}
		// 不选择当前nums[i], 继续下一层
		recur(i + 1) // chosen无变化

		// 选择当前nums[i], 将其加入chosen中, 继续下一层
		chosen = append(chosen, nums[i])
		recur(i + 1) // chosen增加了当前nums[i]

		// 回到上一层状态，将当前层所选择的值pop掉
		chosen = chosen[:len(chosen)-1]
	}

	recur(0)
	return ans
}
