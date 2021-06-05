package main

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	//基于2的幂的做法
	// return n > 0 && n&(n-1) == 0 && (n&0xaaaaaaaa) == 0
	if n == 0 {
		return false
	}
	for n != 1 {
		if n&3 != 0 {
			return false
		}
		n >>= 2
	}
	return true
}

//1111
//1010101010

//100 10000 1000000
// func isPowerOfFour(n int) bool {
// 	t := n
// 	if n == 0 {
// 		return false
// 	}
// 	for t != 0 {
// 		x := t % 4
// 		if x != 0 && t != 1 {
// 			return false
// 		}
// 		t = t / 4
// 	}
// 	return true
// }

// @lc code=end
