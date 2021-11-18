package perm

import (
	"fmt"
	"testing"
)

func Test_permState_Next(t *testing.T) {
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

func Test_permStart_Done(t *testing.T) {
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

func TestStartFrom(t *testing.T) {
	testCases := map[string]struct {
		p    []int
		next []int
	}{
		"start->second":   {[]int{2, 3, 1, 0}, []int{3, 1, 0, 2}},
		"small->large":    {[]int{3, 2, 1, 0}, []int{2, 1, 0, 3}},
		"end not advance": {[]int{1, 2, 0, 3}, []int{1, 2, 0, 3}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			g := StartFrom(tc.p)
			g.Next()
			cur := g.Index()
			for i, k := range tc.next {
				if k != cur[i] {
					t.Errorf("wrong next permutation: %v != next:%v", cur, tc.next)
					break
				}
			}
		})
	}
}
