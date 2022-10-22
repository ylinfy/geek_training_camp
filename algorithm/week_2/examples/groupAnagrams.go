import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var groups = make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		// 将排后的字符串作为key, 所有异构的字符串组成列表作为value
		groups[sorted] = append(groups[sorted], str)
	}

	ans := make([][]string, 0, len(groups))
	for _, v := range groups {
		ans = append(ans, v) // 将所有异构字符串列表组成一个二级字符串数组
	}

	return ans
}

func sortString(str string) string {
	byteStr := []byte(str)
	sort.Slice(byteStr, func(i, j int) bool {
		return byteStr[i] < byteStr[j]
	})
	return string(byteStr)
}
