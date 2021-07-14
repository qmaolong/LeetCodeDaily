package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1818 lang=golang
 *
 * [1818] 绝对差值和
 */

//  51/51 cases passed (176 ms)
// Your runtime beats 79.66 % of golang submissions
// Your memory usage beats 100 % of golang submissions (8.9 MB)

// @lc code=start
func minAbsoluteSumDiff(nums1 []int, nums2 []int) (res int) {
	changeValues := make([]int, len(nums1))

	c := make([]int, len(nums1))
	copy(c, nums1)
	sort.Ints(c)

	max := 0
	changeI := 0
	for i, v := range nums1 {
		nearestValue, _ := findNearestValue(c, nums2[i])
		changeValues[i] = nearestValue
		v = abs(v-nums2[i]) - abs(nearestValue-nums2[i])
		if v > max {
			max = v
			changeI = i
		}
	}
	nums1[changeI] = changeValues[changeI]
	for i, v := range nums1 {
		res += abs(v - nums2[i])
	}
	return res % 1000000007
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findNearestValue(arr []int, target int) (int, int) {
	cursor := len(arr) / 2
	v := arr[cursor]
	distance := abs(target - v)
	if target == v || len(arr) == 1 {
		return v, distance
	}
	if v > target {
		arr = arr[0:cursor]
	} else {
		arr = arr[cursor+1:]
	}
	if len(arr) == 0 {
		return v, distance
	}
	a, b := findNearestValue(arr, target)
	if b < distance {
		return a, b
	} else {
		return v, distance
	}
}

// @lc code=end
