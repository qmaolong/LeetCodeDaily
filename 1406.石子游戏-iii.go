package main

/*
 * @lc app=leetcode.cn id=1406 lang=golang
 *
 * [1406] 石子游戏 III
 */

// @lc code=start
func stoneGameIII(stoneValue []int) string {
	prefixSum := make([]int, len(stoneValue))
	aScore := 0
	bScore := 0
	for i, v := range stoneValue {
		if i == 0 {
			prefixSum[i] = v
		} else {
			prefixSum[i] = prefixSum[i-1] + v
		}
	}

}

func bestMove(arr []int) (int, int) {
	bestMove := 1
	sum := 0

	sum1 := arr[0]
	m := sum1 - maxSum(arr[1:])
	sum = sum1

	sum2 := arr[0] + arr[1]
	step2 := sum2 - maxSum(arr[2:])
	if step2 > m {
		bestMove = 2
		m = step2
		sum = sum2
	}
	sum3 := arr[0] + arr[1] + arr[2]
	step3 := sum3 - maxSum(arr[3:])
	if step3 > m {
		bestMove = 3
		m = step3
		sum = sum3
	}
	return bestMove, sum
}

func maxSum(arr []int) int {
	m := 0
	s := 0
	for i := 0; i < 3 && i < len(arr); i++ {
		s += arr[i]
		if s > m {
			m = s
		}
	}
	return m
}

// @lc code=end

//[1,2,3,4,18]
