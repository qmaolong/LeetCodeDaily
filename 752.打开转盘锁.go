package main

/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */
//  48/48 cases passed (104 ms)
//  Your runtime beats 70.24 % of golang submissions
//  Your memory usage beats 50.24 % of golang submissions (7.8 MB)
// @lc code=start
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	seadMap := make(map[string]bool)
	for _, d := range deadends {
		seadMap[d] = true
	}
	if seadMap["0000"] {
		return -1
	}

	type item struct {
		value string
		step  int
	}
	queue := make([]item, 0)
	queue = append(queue, item{"0000", 0})
	for len(queue) > 0 {
		it := queue[0]
		if it.value == target {
			return it.step
		}
		queue = queue[1:]
		arr := []rune(it.value)
		for i, v := range arr {
			var num = int(v - '0')
			arr[i] = rune((num+1)%10 + '0')
			t1 := string(arr)
			if !seadMap[t1] {
				queue = append(queue, item{t1, it.step + 1})
				seadMap[t1] = true
			}
			if t1 == target {
				return it.step + 1
			}
			arr[i] = rune((num+9)%10 + '0')
			t2 := string(arr)
			if !seadMap[t2] {
				queue = append(queue, item{t2, it.step + 1})
				seadMap[t2] = true
			}
			if t2 == target {
				return it.step + 1
			}
			arr[i] = v
		}
	}
	return -1
}

// @lc code=end
