package perm

import "testing"

const benchNum = 7

func BenchmarkPrimitive(b *testing.B) {
	p := make([]int, benchNum)
	for i := 0; i < b.N; i++ {
		if err := Init(p); err != nil {
			b.Fatalf("Init() failed: %v", err)
		}
		for !IsEnd(p) {
			Advance(p)
		}
	}
}

func BenchmarkGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p, err := New(benchNum)
		if err != nil {
			b.Fatalf("New() failed: %v", err)
		}
		for !p.Done() {
			p.Next()
		}
	}
}

func BenchmarkPermRecursive(b *testing.B) {
	n := benchNum
	p := make([]int, n)
	ignore := make([]bool, n)
	for i := 0; i < b.N; i++ {
		recPerm(p, n, ignore)
	}
}

func recPerm(p []int, n int, ignore []bool) {
	if n == 0 {
		return
	}
	for i := 0; i < len(p); i++ {
		if !ignore[i] {
			p[n-1] = i
			ignore[i] = true
			recPerm(p, n-1, ignore)
			ignore[i] = false
		}
	}
}
