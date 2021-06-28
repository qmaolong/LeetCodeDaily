package main

/*
 * @lc app=leetcode.cn id=909 lang=golang
 *
 * [909] 蛇梯棋
 */
//  211/211 cases passed (20 ms)
//  Your runtime beats 42.86 % of golang submissions
//  Your memory usage beats 21.43 % of golang submissions (6.6 MB)

// @lc code=start
func snakesAndLadders(board [][]int) int {
	l := len(board)
	target := l * l

	link := make([]int, 0)
	i := l - 1
	j := 0
	factor := 1
	for len(link) < target {
		p := board[i][j]
		link = append(link, p)
		if (j == l-1 && factor == 1) || (j == 0 && factor == -1) {
			i--
			factor *= -1
		} else {
			j += factor
		}
	}

	type item struct {
		value int
		step  int
	}

	queue := make([]item, 0)
	queue = append(queue, item{
		value: 1,
		step:  0,
	})

	sead := make(map[int]bool)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		var t int
		for i := 1; i <= 6 && (v.value+i) <= target; i++ {
			t = link[v.value+i-1]
			if v.value+i == target || t == target {
				return v.step + 1
			}
			if t == -1 {
				t = v.value + i
			}
			if !sead[t] {
				sead[t] = true
				queue = append(queue, item{
					value: t,
					step:  v.step + 1,
				})
			}
		}
	}
	return -1
}

// @lc code=end
