package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=726 lang=golang
 *
 * [726] 原子的数量
 */

// 31/31 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 12.12 % of golang submissions (2.7 MB)

// @lc code=start
type Item struct {
	str   string
	count int
	plus  []*int
}

type SortItem []Item

func (a SortItem) Len() int           { return len(a) }
func (a SortItem) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortItem) Less(i, j int) bool { return a[i].str < a[j].str }

func countOfAtoms(formula string) string {
	items := make([]Item, 0)
	cursor := 0
	plusStack := make([]*int, 0)
	for cursor < len(formula) {
		names := []byte{}
		names = append(names, formula[cursor])
		if formula[cursor] == '(' {
			num := 1
			plusStack = append(plusStack, &num)
		} else if formula[cursor] >= 'A' && formula[cursor] <= 'Z' {
			for cursor+1 < len(formula) && formula[cursor+1] >= 'a' && formula[cursor+1] <= 'z' {
				cursor++
				names = append(names, formula[cursor])
			}
			item := Item{
				str:   string(names),
				count: 1,
				plus:  make([]*int, 0),
			}
			item.plus = append(item.plus, plusStack...)
			items = append(items, item)
		} else if formula[cursor] >= '0' && formula[cursor] <= '9' {
			prev := formula[cursor-1]
			for cursor+1 < len(formula) && formula[cursor+1] <= '9' && formula[cursor+1] >= '0' {
				cursor++
				names = append(names, formula[cursor])
			}
			num, _ := strconv.Atoi(string(names))
			if prev == ')' {
				*plusStack[len(plusStack)-1] = num
				plusStack = plusStack[0 : len(plusStack)-1]
			} else {
				items[len(items)-1].count = num
			}
		}
		cursor++
	}
	m := make(map[string]*Item)
	for i := range items {
		v := items[i]
		it := m[v.str]
		for _, v1 := range v.plus {
			v.count *= *v1
		}
		if it == nil {
			m[v.str] = &v
		} else {
			it.count += v.count
		}
	}
	items = make([]Item, 0)
	for _, v := range m {
		items = append(items, *v)
	}
	sort.Sort(SortItem(items))
	res := ""
	for _, v := range items {
		res += v.str
		if v.count > 1 {
			res += strconv.Itoa(v.count)
		}
	}
	return res
}

// @lc code=end
