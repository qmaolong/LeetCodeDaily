/*
 * @lc app=leetcode.cn id=233 lang=java
 *
 * [233] 数字 1 的个数
 */

// @lc code=start
class Solution {
    public int countDigitOne(int n) {
        String s = String.valueOf(n);
        int length = s.length();
        int count = 0;
        for (int i = 0; i < length; i++) {
            int temp = 0;
            int num = n % 10;
            if (num > 0) {
                temp += 1;
            }
            temp += num * temp;
            count += temp;
            n = n/10;
        }
        return count;
    }

    // public static void main(String[] args) {
    //     int res = new Solution233().countDigitOne(13);
    //     System.out.println(res == 6);
    //     res = new Solution233().countDigitOne(0);
    //     System.out.println(res == 0);
    // }
}
// @lc code=end

