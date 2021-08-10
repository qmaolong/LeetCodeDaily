/*
 * @lc app=leetcode.cn id=413 lang=java
 *
 * [413] 等差数列划分
 */
// 15/15 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 80.06 % of java submissions (36.1 MB)

// @lc code=start
class Solution {
    public int numberOfArithmeticSlices(int[] nums) {
        int n = nums.length;
        if(n <= 2) return 0;
        int dp = 0;
        int ans = 0;
        for(int i=3;i<=n;++i){
            if(nums[i-1]-nums[i-2] == nums[i-2]-nums[i-3]) dp++;
            else dp = 0;
            ans += dp;
        }
        return ans;
    }

    public static void main(String[] args) {
        int res = new Solution().numberOfArithmeticSlices(new int[]{1,2,3,4,6,7,8,9});
        System.out.println(res == 6);
    }
}
// @lc code=end

