package main

import "testing"

var tests = []struct {
	nums []int64
	sum  int64
}{
	{nums: gen(1 << 2), sum: 6}, //0 1 2 3 (0+3)/2 * 4
	{nums: gen(1 << 3), sum: 28},
	{nums: gen(1 << 5), sum: 496},
	{nums: gen(1 << 7), sum: 8128},
	{nums: gen(1 << 9), sum: 130_816},
	{nums: gen(1 << 11), sum: 2_096_128},
	{nums: gen(1 << 13), sum: 33_550_336},
	{nums: gen(1 << 18), sum: 34_359_607_296}, //(0+2^18-1)/2*2^18
	{nums: gen(1 << 26), sum: 2_251_799_780_130_816},
}

func TestSumCLS(t *testing.T) {
	for i := range tests {
		test := tests[i]
		res := sumCLS(test.nums)
		if res != test.sum {
			t.Fatalf("expected %d, but got %d", test.sum, res)
		}
	}
}

func TestSumASM(t *testing.T) {
	for i := range tests {
		test := tests[i]
		res := sumASM(test.nums)
		if res != test.sum {
			t.Fatalf("expected %d, but got %d", test.sum, res)
		}
	}
}

func TestSumAVX(t *testing.T) {
	for i := range tests {
		test := tests[i]
		res := sumAVX(test.nums)
		if res != test.sum {
			t.Fatalf("expected %d, but got %d", test.sum, res)
		}
	}
}

var benchmarks = []struct {
	name string
	nums []int64
}{
	{name: "SumOf4", nums: gen(1 << 2)},
	{name: "SumOf8", nums: gen(1 << 3)},
	{name: "SumOf32", nums: gen(1 << 5)},
	{name: "SumOf128", nums: gen(1 << 7)},
	{name: "SumOf512", nums: gen(1 << 9)},
	{name: "SumOf2048", nums: gen(1 << 11)},
	{name: "SumOf8192", nums: gen(1 << 13)},
	{name: "SumOf262144", nums: gen(1 << 18)},
	{name: "SumOf67M", nums: gen(1 << 26)},
}

func BenchmarkSumCLS(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = sumCLS(bm.nums)
			}
		})
	}
}

func BenchmarkSumASM(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = sumASM(bm.nums)
			}
		})
	}
}

func BenchmarkSumAVX(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = sumAVX(bm.nums)
			}
		})
	}
}
