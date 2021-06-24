package main

/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */
//  33/33 cases passed (4 ms)
//  Your runtime beats 88.42 % of golang submissions
//  Your memory usage beats 90.53 % of golang submissions (2.7 MB)

// @lc code=start
func maxPoints(points [][]int) int {
	l := len(points)

	//判断3个点是否为一条直线
	var ifLine = func(p1, p2, p3 []int) bool {
		a1 := p1[0] - p2[0]
		a2 := p1[1] - p2[1]
		b1 := p2[0] - p3[0]
		b2 := p2[1] - p3[1]

		if a1 == 0 || b1 == 0 {
			return a1 == b1 && p1[0] == p2[0] && p1[0] == p3[0]
		}
		if a2 == 0 || b2 == 0 {
			return a2 == b2 && p1[1] == p2[1] && p1[1] == p3[1]
		}
		return a1*b2 == a2*b1
	}

	arr := make([][]*int, l)
	for i := range arr {
		arr[i] = make([]*int, l)
	}

	max := 1
	for i := range points {
		for j := i + 1; j < l; j++ {
			if arr[i][j] != nil {
				continue
			}
			tCount := 2
			arr[i][j] = &tCount
			for k := j + 1; k < l; k++ {
				if ifLine(points[i], points[j], points[k]) {
					tCount++
					arr[i][k] = &tCount
					arr[j][k] = &tCount
				}
			}
			if tCount > max {
				max = tCount
			}
		}
	}

	return max
}

// @lc code=end
