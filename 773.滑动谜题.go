package main

/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 */
//  32/32 cases passed (0 ms)
//  Your runtime beats 100 % of golang submissions
//  Your memory usage beats 13.64 % of golang submissions (4.5 MB)

// @lc code=start
func slidingPuzzle(board [][]int) int {
	var b [2][3]int
	for i, v := range board {
		for j, v1 := range v {
			b[i][j] = v1
		}
	}
	type item struct {
		value [2][3]int
		step  int
		str   string
	}
	var tostr = func(arr [2][3]int) string {
		var res []byte
		for _, v := range arr {
			for _, v1 := range v {
				res = append(res, byte(v1+'0'))
			}
		}
		return string(res)
	}

	queue := make([]item, 0)
	queue = append(queue, item{
		b,
		0,
		tostr(b),
	})

	sead := make(map[string]bool)
	var addSead = func(arr [2][3]int) {
		str := tostr(arr)
		if !sead[str] {
			sead[str] = true
		}
	}
	for len(queue) > 0 {
		b := queue[0]
		if b.str == "123450" {
			return b.step
		}
		addSead(b.value)
		queue = queue[1:]

		zeroX := 0
		zeroY := 0
		for i, v := range b.value {
			for j, v1 := range v {
				if v1 == 0 {
					zeroX = i
					zeroY = j
					break
				}
			}
		}
		if zeroX == 0 {
			//0向下移
			c := b.value
			t := c[zeroX+1][zeroY]
			c[zeroX+1][zeroY] = 0
			c[zeroX][zeroY] = t
			str := tostr(c)
			if str == "123450" {
				return b.step + 1
			}
			if !sead[str] {
				queue = append(queue, item{
					c,
					b.step + 1,
					str,
				})
			}
		} else {
			//向上移
			c := b.value
			t := c[zeroX-1][zeroY]
			c[zeroX-1][zeroY] = 0
			c[zeroX][zeroY] = t
			str := tostr(c)
			if str == "123450" {
				return b.step + 1
			}
			if !sead[str] {
				queue = append(queue, item{
					c,
					b.step + 1,
					str,
				})
			}
		}
		//向左
		if zeroY > 0 {
			c := b.value
			t := c[zeroX][zeroY-1]
			c[zeroX][zeroY-1] = 0
			c[zeroX][zeroY] = t
			str := tostr(c)
			if str == "123450" {
				return b.step + 1
			}
			if !sead[str] {
				queue = append(queue, item{
					c,
					b.step + 1,
					str,
				})
			}
		}
		//向右
		if zeroY < 2 {
			c := b.value
			t := c[zeroX][zeroY+1]
			c[zeroX][zeroY+1] = 0
			c[zeroX][zeroY] = t
			str := tostr(c)
			if str == "123450" {
				return b.step + 1
			}
			if !sead[str] {
				queue = append(queue, item{
					c,
					b.step + 1,
					str,
				})
			}
		}
	}
	return -1
}

// @lc code=end
