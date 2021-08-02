package main

/*
 * @lc app=leetcode.cn id=218 lang=golang
 *
 * [218] 天际线问题
 */

// @lc code=start
func getSkyline(buildings [][]int) [][]int {
	optimized := make([][]int, 0)
	for _, v := range buildings {
		optimizeBuildings(&v, &optimized)
	}
	res := make([][]int, 0)
	// for _, v := range optimized {
	// 	left := []int{v[0], v[2]}
	// 	right := []int{v[1], v[2]}

	// }
	return res
}

func optimizeBuildings(buildings *[]int, optimized *[][]int) {

}

// @lc code=end
