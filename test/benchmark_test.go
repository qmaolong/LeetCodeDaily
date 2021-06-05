package test

import (
	"testing"
)

func BenchmarkOne(b *testing.B) {
	// TODO: Initialize
	nums := []int{2, 4, 14}
	for i := 0; i < b.N; i++ {
		_ = totalHammingDistance(nums)
	}

}

func BenchmarkTwo(b *testing.B) {
	// TODO: Initialize
	nums := []int{2, 4, 14}
	for i := 0; i < b.N; i++ {
		_ = totalHammingDistance2(nums)
	}
}

func BenchmarkThree(b *testing.B) {
	// TODO: Initialize
	nums := []int{2, 4, 14}
	for i := 0; i < b.N; i++ {
		_ = totalHammingDistance3(nums)
	}
}

//耗时114 ns/op
func totalHammingDistance(nums []int) int {
	total := 0
	l := len(nums)
	for i := 0; i < 30; i++ {
		oneCount := 0
		for j, v := range nums {
			tmp := v % 2
			if tmp == 1 {
				oneCount++
			}
			nums[j] = v / 2
		}
		total += oneCount * (l - oneCount)
	}
	return total
}

// 耗时836 ns/op
func totalHammingDistance2(nums []int) int {
	total := 0
	l := len(nums)
	for i := 0; i < 30; i++ {
		oneCount := 0
		a := 2 << i
		b := 1 << i
		for _, v := range nums {
			tmp := v % a
			if tmp >= b { //耗时操作,比tmp==1耗时太多了，也不知道为啥
				oneCount++
			}
		}
		total += oneCount * (l - oneCount)
	}
	return total
}

func totalHammingDistance3(nums []int) int {
	total := 0
	l := len(nums)
	for i := 0; i < 30; i++ {
		oneCount := 0
		for _, v := range nums {
			oneCount += v >> i & 1
		}
		total += oneCount * (l - oneCount)
	}
	return total
}
