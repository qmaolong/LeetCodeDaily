package main

/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */
//  Your runtime beats 88.42 % of golang submissions
//  Your memory usage beats 90.53 % of golang submissions (2.7 MB)

// @lc code=start
func maxPoints(points [][]int) int {

	//判断3个点是否为一条直线
	var ifLine = func(p1, p2, p3 []int) bool {
		if (p1[0]-p2[0]) == 0 || (p2[0]-p3[0]) == 0 {
			return (p2[0]-p3[0]) == (p1[0]-p2[0]) && p1[0] == p2[0] && p1[0] == p3[0]
		}
		if p1[1]-p2[1] == 0 || p2[1]-p3[1] == 0 {
			return p1[1]-p2[1] == p2[1]-p3[1] && p1[1] == p2[1] && p1[1] == p3[1]
		}
		return (p1[0]-p2[0])*(p2[1]-p3[1]) == (p1[1]-p2[1])*(p2[0]-p3[0])
	}

	arr := make([][]*int, len(points))
	for i := range arr {
		arr[i] = make([]*int, len(points))
	}

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			if arr[i][j] != nil && *arr[i][j] > 0 {
				continue
			}
			tCount := 2
			arr[i][j] = &tCount
			for k := j + 1; k < len(points); k++ {
				if ifLine(points[i], points[j], points[k]) {
					tCount++
					arr[i][k] = &tCount
					arr[j][k] = &tCount
				}
			}
		}
	}

	max := 1
	for i := range arr {
		for j := i + 1; j < len(arr[i]); j++ {
			if arr[i][j] != nil && *arr[i][j] > max {
				max = *arr[i][j]
			}
		}
	}
	return max
}

// @lc code=end
