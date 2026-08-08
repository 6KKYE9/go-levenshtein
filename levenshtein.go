package main

// Distance 计算两个字符串之间的莱文斯坦（编辑）距离：
// 把一个变成另一个最少要改多少个字符（增/删/替）。
func Distance(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	n, m := len(ra), len(rb)
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	// 只用两行滚动，省内存
	prev := make([]int, m+1)
	cur := make([]int, m+1)
	for j := 0; j <= m; j++ {
		prev[j] = j
	}
	for i := 1; i <= n; i++ {
		cur[0] = i
		for j := 1; j <= m; j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			del := prev[j] + 1
			ins := cur[j-1] + 1
			sub := prev[j-1] + cost
			best := del
			if ins < best {
				best = ins
			}
			if sub < best {
				best = sub
			}
			cur[j] = best
		}
		prev, cur = cur, prev
	}
	return prev[m]
}
