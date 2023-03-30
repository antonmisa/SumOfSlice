package main

import (
	"fmt"
)

//go:noescape
func _sumASM(nums []int64) int64 // implemented externally

func sumASM(nums []int64) int64 {
	return _sumASM(nums)
}

//go:noescape
func _sumAVX(x []int64) int64 // implemented externally

func sumAVX(nums []int64) int64 {
	return _sumAVX(nums)
}

func sumCLS(nums []int64) int64 {
	var sum int64 = 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

func gen(N int) []int64 {
	nums := make([]int64, N)
	for i := 0; i < N; i++ {
		nums[i] = int64(i)
	}
	return nums
}

func main() {
	// Create an array of integers
	nums := gen(1 << 13)
	fmt.Println(len(nums), cap(nums))

	fmt.Printf("Classic sum: %d\n", sumCLS(nums))
	fmt.Printf("ASM sum: %d\n", sumASM(nums))
	fmt.Printf("AVX sum: %d\n", sumAVX(nums))
}
