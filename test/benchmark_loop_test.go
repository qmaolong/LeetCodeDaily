package test

import "testing"

//BenchmarkLoop1-4   	1000000000	         0.752 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLoop1(b *testing.B) {
	// TODO: Initialize
	a := 2 * b.N //循环次数加倍
	for i := 0; i < a; i++ {
		_ = 100 + 100
	}
}

//BenchmarkLoop2-4   	1000000000	         0.390 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLoop2(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		a := 100 + 100
		b := a + a
		_ = b + b //计算加倍
	}
}

//事实证明循环加倍导致耗时加倍，但是计算加倍影响很小
