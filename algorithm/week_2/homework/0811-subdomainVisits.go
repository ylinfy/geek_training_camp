import (
	"strconv"
	"strings"
)

// 使用哈希表进行计数，将每个子域名访问次数累加起来
func subdomainVisits(cpdomains []string) []string {
	visitsNums := make(map[string]int)

	for _, cpdomain := range cpdomains {
		// 分解cpdomain, 得到访问次数num和最低一级的域名subdomain
		i := strings.IndexByte(cpdomain, ' ')
		num, _ := strconv.Atoi(cpdomain[:i])
		subdomain := cpdomain[i+1:]

		// 循环获取二级域名和顶级域名，并将其对应的访问次数累加起来
		for {
			visitsNums[subdomain] += num
			i := strings.IndexByte(subdomain, '.')
			// 当未找到'.'时，说明当前域名为顶级域名，所有域名均已找出，退出当前循环
			if i < 0 {
				break
			}
			// 更新域名为更高级的域名
			subdomain = subdomain[i+1:]
		}
	}

	// 从map中获取对应的数据并返回
	ans := []string{}
	for k, v := range visitsNums {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}

	return ans
}
