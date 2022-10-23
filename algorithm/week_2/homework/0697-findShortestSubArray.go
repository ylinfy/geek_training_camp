type numInfo struct {
	// 分别代该数值出的次数，出现在列表的最左端下标和最右端下标
	cnt, left, right int
}

// 记录每个数值出现的次数，出现的起始下标位置和最终下标位置
func findShortestSubArray(nums []int) int {
	numsInfos := make(map[int]numInfo)

	for i, num := range nums {
		// 如果已经存在该数值，更新其计数和最终下标位置
		if info, ok := numsInfos[num]; ok {
			info.cnt++
			info.right = i
			numsInfos[num] = info
		} else {
			// 如果是首次遇到的数值，直接存储起来, 起始和最终下标均初始化为当前下标
			numsInfos[num] = numInfo{1, i, i}
		}
	}

	// 遍历map，寻找最大频次数及对应数值的最短子数组
	maxCnt, ans := 0, 0
	for _, info := range numsInfos {
		// 频次过大，更新maxCnt和其下标数据
		if info.cnt > maxCnt {
			maxCnt, ans = info.cnt, info.right-info.left+1
		} else if info.cnt == maxCnt { // 频次相等，比较下标跨度更小
			if info.right-info.left+1 < ans {
				ans = info.right - info.left + 1
			}
		}
	}

	return ans
}
