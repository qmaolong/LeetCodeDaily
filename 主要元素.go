package main

//面试题 17.10. 主要元素
// 执行用时：20 ms, 在所有 Go 提交中击败了80.49%的用户
// 内存消耗：6 MB, 在所有 Go 提交中击败了78.86%的用户

func majorityElement(nums []int) (res int) {
	count := 0
	for _, v := range nums {
		if v == res {
			count++
		} else if count > 0 {
			count--
		} else {
			res = v
			count = 1
		}
	}
	count = 0
	for _, v := range nums {
		if v == res {
			count++
		}
	}
	if count <= len(nums)/2 {
		return -1
	}
	return
}
