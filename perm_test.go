package perm

import (
	"fmt"
	"testing"
)

func TestPerm_Next(t *testing.T) {
	testcase := []struct {
		n   int
		ns  []int
		all int
	}{
		{0, nil, 0},
		{1, []int{1}, 0},
		{2, []int{1, 2}, 1 * 2},
		{3, []int{1, 2, 3}, 1 * 2 * 3},
		{4, []int{1, 2, 3, 4}, 1 * 2 * 3 * 4},
		{5, []int{1, 2, 3, 4, 5}, 1 * 2 * 3 * 4 * 5},
	}
	for _, tc := range testcase {
		table := make(map[string]struct{}, tc.all)

		p := Iter(tc.ns)
		for i := 0; i < tc.all; i++ {
			s := fmt.Sprintf("%v", tc.ns)
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
		ns  []int
		all int
	}{
		{0, nil, 0},
		{1, []int{1}, 0},
		{2, []int{1, 2}, 1 * 2},
		{3, []int{1, 2, 3}, 1 * 2 * 3},
		{4, []int{1, 2, 3, 4}, 1 * 2 * 3 * 4},
		{5, []int{1, 2, 3, 4, 5}, 1 * 2 * 3 * 4 * 5},
	}
	for _, tc := range testcase {
		p := Iter(tc.ns)
		if tc.n > 1 {
			for i := 0; i < tc.all-1; i++ {
				if p.Done() {
					t.Errorf("n=%d: not finished: Done got true, want false: [%d] %v", tc.n, i, tc.ns)
				}
				p.Next()
			}
		}

		if tc.n > 1 && p.Done() {
			t.Errorf("n=%d: not finished: Done got true, want false: %v", tc.n, tc.ns)
		}
		p.Next()
		if !p.Done() {
			t.Errorf("n=%d: finished: Done got false, want true: %v", tc.n, tc.ns)
		}
		p.Next()
		if !p.Done() {
			t.Errorf("n=%d: finished: Done got still true, want false: %v", tc.n, tc.ns)
		}
	}
}

func TestPerm_StartFrom(t *testing.T) {
	{
		small := []int{2, 3, 1, 0}
		p := StartFrom(small, nil)
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

const benchNum = 11

func BenchmarkNext(b *testing.B) {
	p := New(benchNum)
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
