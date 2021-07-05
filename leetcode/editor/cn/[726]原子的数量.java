//给定一个化学式formula（作为字符串），返回每种原子的数量。 
//
// 原子总是以一个大写字母开始，接着跟随0个或任意个小写字母，表示原子的名字。 
//
// 如果数量大于 1，原子后会跟着数字表示原子的数量。如果数量等于 1 则不会跟数字。例如，H2O 和 H2O2 是可行的，但 H1O2 这个表达是不可行的。
// 
//
// 两个化学式连在一起是新的化学式。例如 H2O2He3Mg4 也是化学式。 
//
// 一个括号中的化学式和数字（可选择性添加）也是化学式。例如 (H2O2) 和 (H2O2)3 是化学式。 
//
// 给定一个化学式，输出所有原子的数量。格式为：第一个（按字典序）原子的名子，跟着它的数量（如果数量大于 1），然后是第二个原子的名字（按字典序），跟着它的数
//量（如果数量大于 1），以此类推。 
//
// 示例 1: 
//
// 
//输入: 
//formula = "H2O"
//输出: "H2O"
//解释: 
//原子的数量是 {'H': 2, 'O': 1}。
// 
//
// 示例 2: 
//
// 
//输入: 
//formula = "Mg(OH)2"
//输出: "H2MgO2"
//解释: 
//原子的数量是 {'H': 2, 'Mg': 1, 'O': 2}。
// 
//
// 示例 3: 
//
// 
//输入: 
//formula = "K4(ON(SO3)2)2"
//输出: "K4N2O14S4"
//解释: 
//原子的数量是 {'K': 4, 'N': 2, 'O': 14, 'S': 4}。
// 
//
// 注意: 
//
// 
// 所有原子的第一个字母为大写，剩余字母都是小写。 
// formula的长度在[1, 1000]之间。 
// formula只包含字母、数字和圆括号，并且题目中给定的是合法的化学式。 
// 
// Related Topics 栈 哈希表 字符串 
// 👍 165 👎 0


import com.alibaba.zhengjuzx.controller.console.Test;
import org.apache.commons.lang.StringUtils;

import java.util.*;
import java.util.concurrent.LinkedBlockingDeque;
import java.util.stream.Collectors;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public String countOfAtoms(String formula) {
        List<String> list = new ArrayList<>();
        while (formula.length() > 0){
            String tmp = formula.replaceFirst("[0-9]+|[A-Z][a-z]*|\\(|\\)", "");
            list.add(formula.substring(0, formula.length() - tmp.length()));
            formula = tmp;
        }
        Deque<Test.Item> stack = new LinkedBlockingDeque<>();
        for (int i = 0; i < list.size(); i++) {
            String str = list.get(i);
            String next = i == list.size() - 1? null: list.get(i+1);
            int num = 1;
            if (StringUtils.isNumeric(next)){
                num = Integer.parseInt(next);
            }
            if (")".equals(str)){
                List<Test.Item> tempList = new ArrayList();
                while (true){
                    Test.Item item = stack.poll();
                    if ("(".equals(item.str)){
                        break;
                    }
                    item.count *= num;
                    tempList.add(item);
                }
                for (int j = tempList.size() - 1; j >= 0; j--) {
                    stack.push(tempList.get(j));
                }
            }else if (!StringUtils.isNumeric(str)){
                stack.push(new Test.Item(str, num));
            }
        }
        Map<String, Integer> map = new HashMap<>();
        for (Test.Item item : stack) {
            map.merge(item.str, item.count, Integer::sum);
        }
        map.entrySet().stream().sorted(Map.Entry.comparingByKey()).map(entry -> entry.getKey() + (entry.getValue() == 1?"":entry.getValue())).collect(Collectors.toList());
    }
}
//leetcode submit region end(Prohibit modification and deletion)
