package main

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240]  搜索二维矩阵 II
 */

//评测题目: 无
//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的是否存在目标值 target 。该矩阵具有以下特性：

//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。

//输入：matrix = [
//	[1,  4,  7,  11,  15, 16],
//	[2,  5,  8,  12,  19, 20],
//	[3,  6,  9,  16,  22, 25],
//	[10, 13, 14, 17,  24, 27],
//	[18, 21, 23, 26,  30, 35]
//], target = 5
//输出：true

//更优解

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	i := m - 1
	j := 0
	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}

// @lc code=end

func match(matrix [][]int, target int, x, y int) bool {
	if matrix[x][y] > target {
		return false
	}
	//记录开始位置
	xt := x
	yt := y
	for x < len(matrix) && y < len(matrix[0]) {
		cursor := matrix[x][y]
		//fmt.Println(cursor)
		if cursor == target {
			return true
		} else if cursor > target {
			//找到大于目标值的元素，
			//以该元素最上方的元素和最左边的元素为开始元素递归查找
			m := match(matrix, target, xt, y)
			if m {
				return m
			}
			m = match(matrix, target, x, yt)
			if m {
				return m
			}
		} else if x == len(matrix)-1 && y != len(matrix[0])-1 {
			//x到达底部，y还没有结束
			return match(matrix, target, xt, y+1)
		} else if y == len(matrix[0])-1 && x != len(matrix)-1 {
			//y到达最右边，x还没有结束
			return match(matrix, target, x+1, yt)
		}
		x++
		y++
	}
	return false
}
