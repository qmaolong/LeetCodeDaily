import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=68 lang=java
 *
 * [68] 文本左右对齐
 */

// 执行用时：1 ms, 在所有 Java 提交中击败了51.71%的用户
// 内存消耗：37 MB, 在所有 Java 提交中击败了20.24%的用户
// 通过测试用例：27 / 27

// @lc code=start
class Solution68 {
    public List<String> fullJustify(String[] words, int maxWidth) {
        List<String> res = new ArrayList<>();
        List<String> line = new ArrayList<>();
        int lineCharCount = 0;
        for (String word : words) {
            boolean enough = lineCharCount + line.size() + word.length() <= maxWidth;
            if (enough) {
                line.add(word);
                lineCharCount += word.length();
            }else {
                int leftBlank = maxWidth - lineCharCount;
                int blankAvg = line.size() > 1? leftBlank / (line.size() - 1): leftBlank;
                int blankRemainder = line.size() > 1? leftBlank % (line.size() - 1): 0;
                StringBuffer buffer = new StringBuffer();
                for (int i = 0; i < line.size(); i++) {
                    String w = line.get(i);
                    buffer.append(w);
                    if (i == line.size() - 1){
                        break;
                    }
                    for (int j = 0; j < blankAvg; j++) {
                        buffer.append(" ");
                    }
                    if (blankRemainder > 0){
                        buffer.append(" ");
                        blankRemainder --;
                    }
                }
                if (line.size() == 1){
                    for (int i = 0; i < maxWidth - lineCharCount; i++) {
                        buffer.append(" ");
                    }
                }
                res.add(buffer.toString());
                line = new ArrayList<>();
                line.add(word);
                lineCharCount = word.length();
            }
        }
        if (line.size() > 0){
            StringBuffer buffer = new StringBuffer();
            buffer.append(String.join(" ", line));
            for (int i = 0; i < maxWidth - lineCharCount - line.size() + 1; i++) {
                buffer.append(" ");
            }
            res.add(buffer.toString());
        }
        return res;
    }

    public static void main(String[] args) {
        String[] words = new String[]{"What","must","be","acknowledgment","shall","be"};
        List<String> res = new Solution68().fullJustify(words, 16);
        System.out.println(res);
    }
}
// @lc code=end

