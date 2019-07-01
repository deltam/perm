package perm

import (
	"fmt"
	"testing"
)

func TestPerm_Next(t *testing.T) {
	ns := []int{1, 2, 3, 4}
	all := 1 * 2 * 3 * 4
	table := make(map[string]struct{}, all)

	p := Iter(ns)
	for !p.Done() {
		s := fmt.Sprintf("%v", ns)
		if _, exist := table[s]; exist {
			t.Errorf("duplicate permutation: %s", s)
		}
		table[s] = struct{}{}
		p.Next()
	}

	if len(table) != all {
		t.Errorf("all permutation count failed: %d", len(table))
	}
}

func TestPerm_Done(t *testing.T) {
	ns := []int{1, 2, 3, 4}
	all := 1 * 2 * 3 * 4
	p := Iter(ns)
	for i := 0; i < all-1; i++ {
		if p.Done() {
			t.Errorf("Done is not done: [%d] %v", i, ns)
		}
		p.Next()
	}

	if p.Done() {
		t.Errorf("Done is not done: %v", ns)
	}
	p.Next()
	if !p.Done() {
		t.Errorf("Done is done: %v", ns)
	}
	p.Next()
	if !p.Done() {
		t.Errorf("Done is keeping done: %v", ns)
	}
}
