```
解答成功:
			执行耗时:4 ms,击败了97.14% 的Java用户
			内存消耗:37 MB,击败了56.43% 的Java用户
```

# 小细节：
```
       写需要自己解析输入（字符串）的题目，千万别头铁将代码都写一起，会疯的~~~。
       我自己的处理方式就是能抽成方法的就抽离为方法，不行就要想抽成独立的类。
       这样主干代码才会清晰，好Debug~~
```

这里是定义了一个专门用来解析输入的化学式的类：FormulaReader

# 主要思路：
      循环体：
          1、判断当前元素是不是 '(' ？ 是则压栈，注意这里压的是atomList的当前大小，本趟循环结束。
          2、判断当前元素是不是 ')' ？
              是。进行下面操作后，结束本趟循环。
              2.1 读取')'后面的数字 count。
              2.2 弹栈获得一个下标值（atomList的）。
              2.3 从下标值开始的化学元素的个数都要乘以count
          3、读取一个化学元素，读取化学元素后面的数字，将化学元素与个数保存到atomList里面
      汇总：
          使用 TreeMap 汇总 atomList里面的化学元素。


# 代码
```java
class Solution {
    //726
    public String countOfAtoms(String formula) {
        // "(" 右边第一个化学元素下标,括号会影响的第一个开始元素，其实就是atomList.size() 哈哈
        // K4(ON(SO3)2)2  出现的第一个'(' 右边的第一个元素是 O，这个 O 在 atomList 里面的小标是 1
        // 所以 将 1 压栈。
        // 这里压下标的操作，是改进了好几次才想到的 ~《-.-》~
        Deque<Integer> stack = new ArrayDeque<>();
        List<Atom> atomList = new ArrayList<>();

        //定义专门用来读化学式的类（职责分离，哈哈^_^）
        FormulaReader reader = new FormulaReader(formula);
        while (reader.hasNext()) {
            //左括号，压栈
            if (reader.curChar() == '(') {
                stack.push(atomList.size());
                reader.index++;
                //这里一定要结束本趟循环
                continue;
            }

            //右括号，弹栈 + 获取右括号后面的数字 + 计算括号内的化学元素
            if (reader.curChar() == ')') {
                reader.index++;
                int count = reader.getCount();
                count = count == 0 ? 1 : count;

                //括号会影响到的化学元素
                int affectIndex = stack.pop();
                while (affectIndex < atomList.size()) {
                    Atom atom = atomList.get(affectIndex);
                    atom.cnt *= count;
                    affectIndex++;
                }
                //这里一定要结束本趟循环
                continue;
            }

            //读化学元素
            String atom = reader.getAtom();
            //读化学元素后面的数字
            int count = reader.getCount();
            count = count == 0 ? 1 : count;
            atomList.add(new Atom(atom, count));
        }

        //汇总元素个数。因为atomList里面的化学元素是有重复的（但数量不一样）
        Map<String, Integer> map = new TreeMap<>();
        for (Atom atom : atomList) {
            int cnt = map.getOrDefault(atom.atom, 0);
            map.put(atom.atom, cnt + atom.cnt);
        }
        StringBuilder builder = new StringBuilder();
        map.forEach((atom, cnt) -> {
            builder.append(atom);
            if (cnt > 1) builder.append(cnt);
        });
        return builder.toString();
    }

    static class FormulaReader {
        String formula;
        int index = 0;

        public FormulaReader(String formula) {
            this.formula = formula;
        }

        public boolean hasNext() {
            return index < formula.length();
        }

        public char curChar() {
            return index < formula.length() ? formula.charAt(index) : '#';
        }

        public int getCount() {
            int cnt = 0;
            char c = curChar();
            while ('0' <= c && c <= '9') {
                cnt = cnt * 10 + c - '0';
                index++;
                c = curChar();
            }
            return cnt;
        }

        public String getAtom() {
            StringBuilder atomBuilder = new StringBuilder();
            atomBuilder.append(formula.charAt(index++));
            char c = curChar();
            while ('a' <= c && c <= 'z') {
                atomBuilder.append(c);
                index++;
                c = curChar();
            }
            return atomBuilder.toString();
        }
    }
    static class Atom {
        String atom;
        int cnt;

        public Atom(String atom, int cnt) {
            this.atom = atom;
            this.cnt = cnt;
        }
    }
}
```

