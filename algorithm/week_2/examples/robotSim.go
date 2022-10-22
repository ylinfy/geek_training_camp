func robotSim(commands []int, obstacles [][]int) int {
	obsSet := make(map[int]struct{})
	for _, obs := range obstacles {
		obsSet[calcHash(obs)] = struct{}{}
	}

	// dir: 0->North, 1->East, 2->South, 3->West
	// turn right: dir = (dir + 1) % 4, 比如当前方向是North,值0，向右计算得到1，即向东
	// turn left: dir = (dir + 3) % 4, 当前方向为South, 向左转，将是向东，同理计算 (2 + 3) % 4 == 1, 即向东
	// 网格中行走技巧: 方向数组, 下一步坐标：(x + dx[dir], y + dy[dir])
	var dx = []int{0, 1, 0, -1}
	var dy = []int{1, 0, -1, 0}

	x, y, dir, ans := 0, 0, 0, 0
	for _, cmd := range commands {
		if cmd == -1 {
			dir = (dir + 1) % 4 // turn right
			continue
		}
		if cmd == -2 {
			dir = (dir + 3) % 4 // turn left
			continue
		}
		for i := 0; i < cmd; i++ {
			// 根据当前方向，判断下一步的坐标
			nextX, nextY := x+dx[dir], y+dy[dir]
			nextPos := []int{nextX, nextY}
			// 查看下一步坐标是否为障碍点，如果是障碍点直接停在当前位置
			if _, ok := obsSet[calcHash(nextPos)]; ok {
				break // 退出当前行走路线，停在当前x, y处，并进行下一次的command
			}
			x, y = nextX, nextY
		}
		// 每一次欧式距离最值会出现在直线行走的起点或终点，只需要与最终的坐标进行比较即可
		ans = maxInt(ans, x*x+y*y)
	}
	return ans
}

func calcHash(obs []int) int {
	// 障碍的坐标点落在[-30000, 30000]中，每个坐标点可以用以下方法得到一个唯一的key
	return (obs[0]+30000)*60001 + (obs[1] + 30000)
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
