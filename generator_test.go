package perm

import (
	"fmt"
	"testing"
)

func TestPerm_Next(t *testing.T) {
	testcase := []struct {
		n   int
		all int
	}{
		{0, 0},
		{1, 0},
		{2, 1 * 2},
		{3, 1 * 2 * 3},
		{4, 1 * 2 * 3 * 4},
		{5, 1 * 2 * 3 * 4 * 5},
	}
	for _, tc := range testcase {
		table := make(map[string]struct{}, tc.all)

		p, err := New(tc.n)
		if tc.n < 3 {
			if err == nil {
				t.Errorf("New() must be fail: n=%d", tc.n)
			}
			continue
		}
		if err != nil {
			t.Errorf("New() failed: n=%d, %v", tc.n, err)
			continue
		}

		for i := 0; i < tc.all; i++ {
			s := fmt.Sprintf("%v", p.Index())
			if _, exist := table[s]; exist {
				t.Errorf("n=%d: duplicate permutation: %s", tc.n, s)
			}
			table[s] = struct{}{}
			p.Next()
		}

		if len(table) != tc.all {
			t.Errorf("n=%d: all permutation count failed: %d", tc.n, len(table))
		}
	}
}

func TestPerm_Done(t *testing.T) {
	testcase := []struct {
		n   int
		all int
	}{
		{0, 0},
		{1, 0},
		{2, 1 * 2},
		{3, 1 * 2 * 3},
		{4, 1 * 2 * 3 * 4},
		{5, 1 * 2 * 3 * 4 * 5},
	}
	for _, tc := range testcase {
		p, err := New(tc.n)
		if tc.n < 3 {
			if err == nil {
				t.Errorf("New() must be fail: n=%d", tc.n)
			}
			continue
		}
		if err != nil {
			t.Errorf("New() failed: n=%d, %v", tc.n, err)
			continue
		}
		for i := 1; i <= tc.all-1; i++ {
			if p.Done() {
				t.Errorf("n=%d: not finished: Done got false, want true: [%d] %v", tc.n, i, p.Index())
			}
			p.Next()
		}

		if p.Done() {
			t.Errorf("n=%d: not finished: Done got false, want true: %v", tc.n, p.Index())
		}
		p.Next()
		if !p.Done() {
			t.Errorf("n=%d: finished: Done got true, want false: %v", tc.n, p.Index())
		}
		p.Next()
		if !p.Done() {
			t.Errorf("n=%d: finished: Done got still false, want true: %v", tc.n, p.Index())
		}
	}
}

/*
func TestPerm_StartFrom(t *testing.T) {
	{
		small := []int{2, 3, 1, 0}
		p := StartFrom(small)
		if p.largeCycle {
			t.Errorf("%v is small cycle", small)
		}
	}
	{
		small := []int{3, 1, 0, 2}
		p := StartFrom(small, nil)
		if p.largeCycle {
			t.Errorf("%v is small cycle", small)
		}
	}
	{
		large := []int{1, 2, 0, 3}
		p := StartFrom(large, nil)
		if !p.largeCycle {
			t.Errorf("%v is large cycle", large)
		}
	}
	{
		large := []int{2, 1, 0, 3}
		p := StartFrom(large, nil)
		if !p.largeCycle {
			t.Errorf("%v is large cycle", large)
		}
	}
}
*/

const benchNum = 10

func BenchmarkAdvance(b *testing.B) {
	p := make([]int, benchNum)
	if err := Init(p); err != nil {
		b.Fatalf("Init() failed: %v", err)
	}
	for !IsEnd(p) {
		Advance(p)
	}
}

func BenchmarkNext(b *testing.B) {
	p, err := New(benchNum)
	if err != nil {
		b.Fatalf("New() failed: %v", err)
	}
	for !p.Done() {
		p.Next()
	}
}

func BenchmarkPermRecursive(b *testing.B) {
	n := benchNum
	p := make([]int, n)
	ignore := make([]bool, n)
	recPerm(p, n, ignore)
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
