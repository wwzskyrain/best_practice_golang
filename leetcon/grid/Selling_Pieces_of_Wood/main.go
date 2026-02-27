package main

// 2312. Selling Pieces of Wood 这个题目其实不算难，因为切分只能是从大到小，而且是一刀下去的
func sellingWood(m int, n int, prices [][]int) int64 {
	value := make(map[[2]int]int, 0)
	memo := make([][]int64, m+1)
	for i := range memo {
		memo[i] = make([]int64, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	for _, price := range prices {
		value[[2]int{price[0], price[1]}] = price[2]
	}
	var dfs func(int, int) int64
	dfs = func(x int, y int) int64 {
		if memo[x][y] != -1 {
			return memo[x][y]
		}
		var ret int64
		if val, ok := value[[2]int{x, y}]; ok {
			ret = int64(val)
		}
		if x > 1 {
			for i := 1; i < x; i++ {
				ret = max(ret, dfs(i, y)+dfs(x-i, y))
			}
		}
		if y > 1 {
			for i := 1; i < y; i++ {
				ret = max(ret, dfs(x, i)+dfs(x, y-i))
			}
		}
		memo[x][y] = ret
		return ret
	}
	return dfs(m, n)
}
