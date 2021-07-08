package test

import "testing"

func BenchmarkBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 100
		b := a * 16
		_ = b / 16
	}
}

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 100
		b := a << 4
		_ = b >> 4
	}
}
