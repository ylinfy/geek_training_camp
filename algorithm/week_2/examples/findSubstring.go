var wordCount map[string]int

func findSubstring(s string, words []string) []int {
	// 初始化，分配内存地址
	wordCount = make(map[string]int)
	allLen := 0
	for _, w := range words {
		allLen += len(w) // 记录所有词串起来的总长度
		wordCount[w]++   // 记录words各词数量
	}

	var ans []int // 定义返回的数组（为何不用分配内存）
	for i := 0; i+allLen <= len(s); i++ {
		// 找到满足题意的子串，并记录子串起始index
		if isValid(s[i:i+allLen], words) {
			ans = append(ans, i)
		}
	}
	return ans
}

func isValid(str string, words []string) bool {
	wLen := len(words[0])                 // words中每个词长度一致
	splitWordsMap := make(map[string]int) // 记录str中各长度为wLen的子串数量
	for i := 0; i < len(str); i += wLen {
		splitWordsMap[str[i:i+wLen]]++
	}
	// 比较两个map，如果相同，即该子串str是满足题意的
	return equalsMap(splitWordsMap, wordCount)
}

func equalsMap(m1, m2 map[string]int) bool {
	for k, v := range m1 {
		v2, ok := m2[k]
		if !ok || v2 != v {
			return false
		}
	}
	for k, v := range m2 {
		v2, ok := m1[k]
		if !ok || v2 != v {
			return false
		}
	}
	return true
}
