// copyright author MISUMI Masaru(deltam)

// Package perm provides a permutation generator based on group theory.
package perm

import (
	"reflect"
)

// Perm represents current permutation
type Perm struct {
	cur   []int
	done  bool
	slice interface{}
}

// New returns permutation generator
func New(n int) Perm {
	start := startPerm(n)
	return Perm{cur: start}
}

// Iter generates slice's permutation generator
func Iter(slice interface{}) Perm {
	rv := reflect.ValueOf(slice)
	len := rv.Len()
	start := startPerm(len)
	return Perm{cur: start, slice: slice}
}

// Index returns current permutation as array index
func (p Perm) Index() []int {
	idx := make([]int, len(p.cur))
	copy(idx, p.cur)
	return idx
}

// Next changes order as next permutation
func (p *Perm) Next() {
	if p.done {
		return
	}
	if p.isLast() {
		p.done = true
		return
	}
	if check(p.cur) {
		tau(p.cur)
		tauSlice(p.slice)
	} else {
		sigma(p.cur)
		sigmaSlice(p.slice)
	}
}

// Done returns permutation is all
func (p Perm) Done() bool {
	return p.done
}

func startPerm(n int) []int {
	start := make([]int, n)
	start[0] = n - 2
	start[1] = n - 1
	for i := 2; i < n; i++ {
		start[i] = n - i - 1
	}
	return start
}

func (p Perm) isLast() bool {
	n := len(p.cur)
	if p.cur[n-1] != n-1 {
		return false
	}
	for i := 0; i < n-1; i++ {
		j := i
		if i == 0 {
			j = 1
		} else if i == 1 {
			j = 0
		}
		if p.cur[i] != n-j-2 {
			return false
		}
	}
	return true
}

// 1 2 3 4 ... n-2 n-1 0
func sigma(p []int) {
	n := len(p)
	f := p[0]
	copy(p[0:n-1], p[1:n])
	p[n-1] = f
}

func tau(p []int) {
	p[1], p[0] = p[0], p[1]
}

func sigmaSlice(slice interface{}) {
	if slice == nil {
		return
	}
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	n := rv.Len()
	for i := 1; i < n; i++ {
		swap(i-1, i)
	}
}

func tauSlice(slice interface{}) {
	if slice == nil {
		return
	}
	swap := reflect.Swapper(slice)
	swap(0, 1)
}

func check(p []int) bool {
	n := len(p)
	if p[1] == n-1 {
		return false
	}
	pos := 0
	for i := 0; i < n; i++ {
		if i == 1 {
			continue
		}
		if p[i] == n-1 {
			pos = i
			break
		}
	}
	pos = (pos + 1) % n
	if pos == 1 {
		pos = 2
	}
	if p[pos%n] != (p[1]-2+n)%(n-1) {
		return false
	}
	for i := 0; i < n; i++ {
		if p[i] != n-i-1 {
			return true
		}
	}
	return false
}
