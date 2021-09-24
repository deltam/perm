package perm

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	testCases := map[int][]int{
		3: {1, 2, 0},
		4: {2, 3, 1, 0},
		5: {3, 4, 2, 1, 0},
	}

	for n, init := range testCases {
		p := make([]int, n)
		Init(p)
		for i := range p {
			if init[i] != p[i] {
				t.Errorf("[n=%d]: Init got %v, wants %v", n, p, init)
				break
			}
		}
	}
}

func TestAdvance(t *testing.T) {
	numPerm := 2
	for n := 3; n < 7; n++ {
		numPerm *= n
		exists := make(map[string]struct{}, numPerm)

		p := make([]int, n)
		Init(p)

		for i := 0; i < numPerm; i++ {
			s := fmt.Sprintf("%v", p)
			if _, exist := exists[s]; exist {
				t.Errorf("n=%d: duplicate permutation: %s", n, s)
			}
			exists[s] = struct{}{}
			Advance(p)
		}

		if len(exists) != numPerm {
			t.Errorf("n=%d: all permutation count failed: %d", n, len(exists))
		}
	}
}

func TestSmallCycleIndex(t *testing.T) {
	testCases := []struct {
		pos int
		p   []int
	}{
		{-1, []int{0, 1, 2, 3}},
		{0, []int{2, 3, 1, 0}},
		{1, []int{3, 1, 0, 2}},
		{2, []int{1, 3, 0, 2}},
		{3, []int{3, 0, 2, 1}},
		{2*(4-1) - 1, []int{3, 2, 1, 0}},
		{0, []int{3, 4, 2, 1, 0}},
		{1, []int{4, 2, 1, 0, 3}},
		{2, []int{2, 4, 1, 0, 3}},
		{3, []int{4, 1, 0, 3, 2}},
		{2*(5-1) - 1, []int{4, 3, 2, 1, 0}},
		{-1, []int{4, 0, 2, 3, 1}},
	}

	for i, tc := range testCases {
		idx := SmallCycleIndex(tc.p)
		if idx != tc.pos {
			t.Errorf("[%d]: SmallCycleIndex got %d, wants %d", i+1, idx, tc.pos)
		}
	}
}
