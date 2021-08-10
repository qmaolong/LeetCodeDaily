package main

/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */
// 15/15 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.2 MB)

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	res := 0
	c1 := 0
	c2 := 2
	for c2 < len(nums) {
		if dis := nums[c2] - nums[c2-1]; dis == nums[c1+1]-nums[c1] {
			if c2+1 < len(nums) && nums[c2+1]-nums[c2] == dis {
				c2++
			} else {
				n := c2 - c1 - 1
				res += (n + 1) * n / 2
				c1 = c2 - 1
				c2 = c1 + 2
			}
		} else {
			c1++
			c2++
		}
	}
	return res
}

// @lc code=end

//1 3 6 10 15
//1+1*1/2
//2+1 * 2/2
//3+1 * 3/2
//4+1 * 4/2
//5+1 * 5/2
