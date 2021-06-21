package main

import "regexp"

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
func isNumber(s string) bool {
	r := regexp.MustCompile(`^(\+|\-)?[0-9]*(\.)?([0-9]+)?((e|E)(\+|\-)?[0-9]+)?$`)
	return r.MatchString(s)
}

// @lc code=end
