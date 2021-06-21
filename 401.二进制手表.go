package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

// @lc code=start
func readBinaryWatch(turnedOn int) []string {
	if turnedOn >= 10 {
		return []string{}
	}
	var countLight = func(hour int) int {
		count := 0
		for hour != 0 {
			m := hour & 1
			if m == 1 {
				count++
			}
			hour >>= 1
		}
		return count
	}
	hourMap := make(map[int][]int, 4)
	for i := 0; i < 12; i++ {
		count := countLight(i)
		if hourMap[count] == nil {
			hourMap[count] = make([]int, 0)
		}
		hourMap[count] = append(hourMap[count], i)
	}
	minutMap := make(map[int][]int, 6)
	for i := 0; i < 60; i++ {
		count := countLight(i)
		if minutMap[count] == nil {
			minutMap[count] = make([]int, 0)
		}
		minutMap[count] = append(minutMap[count], i)
	}

	var res []string
	for i := 0; i <= turnedOn && i <= 4; i++ {
		h := hourMap[i]
		m := minutMap[turnedOn-i]
		if len(h) == 0 || len(m) == 0 {
			continue
		}
		for _, hItem := range h {
			for _, mItem := range m {
				mStr := strconv.Itoa(mItem)
				if mItem < 10 {
					mStr = "0" + mStr
				}
				res = append(res, strconv.Itoa(hItem)+":"+mStr)

			}
		}
	}
	return res
}

// @lc code=end
