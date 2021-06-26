package test

import (
	"fmt"
	"testing"
)

func test() {
	a := []int{1, 2, 3}
	b := make([]int, 3)
	c := a
	copy(b, a)
	b[0] = 10
	c[1] = 11

	aa := [][3]int{{1, 2, 3}, {4, 5, 6}}
	bb := make([][3]int, 2)
	copy(bb, aa)
	bb[0][0] = 10
	fmt.Print(b)

}

func Test_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "t1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test()
		})
	}
}
