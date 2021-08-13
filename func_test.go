package perm

import (
	"testing"
)

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
