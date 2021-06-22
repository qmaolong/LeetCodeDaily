package leecode

//输入一个字符串，打印出该字符串中字符的所有排列。
//
// 
//
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。 
//
// 
//
// 示例: 
//
// 输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
// 
//
// 
//
// 限制： 
//
// 1 <= s 的长度 <= 8 
// Related Topics 回溯算法 
// 👍 327 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func permutation(s string) []string {
	arr := []byte(s)
	res := make(map[string]int)
	used := make(map[int]int, 0)
	usedNum := make([]int, 0)

	dfs(arr, used, usedNum, res)
	all := make([]string, 0)
	for k := range res {
		all = append(all, k)
	}
	return all
}

func dfs(arr []byte, used map[int]int, usedNum []int, res map[string]int) {
	if len(arr) == len(used) {
		t := make([]byte, 0)
		for _, v := range usedNum{
			t = append(t, arr[v])
		}
		res[string(t)] = 1
		return
	}
	for i := range arr{
		if used[i] != 0 {
			continue
		}
		copyMap := make(map[int]int)
		for k := range used{
			copyMap[k] = 1
		}
		copySlice := make([]int, len(usedNum))
		copy(copySlice, usedNum)
		copyMap[i] = 1
		copySlice = append(copySlice, i)
		dfs(arr, copyMap, copySlice, res)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
