package test

import (
	"testing"
)

func Test_pointer(t *testing.T) {
	s := make([]*int, 0)
	m := make(map[int]bool)
	n := 1
	n2 := 2
	print(&n)
	print("--->")
	println(n)
	test1(&n)
	println(&n2)
	s = append(s, &n, &n2)
	m[1] = true
	println(&s[0])
	println(&s[1])
	for _, v := range s {
		print(&v)
		print("===>")
		println(v)
		*v = 5
	}
	v1 := 1
	v2 := 2
	var p1 *int
	p1 = &v1
	p1 = &v2
	v1 = 11
	println(p1)
	println("here")
}

func test1(n *int) {
	print(&n)
	print("-->")
	print(n)
	print("-->")
	println(*n)
}
