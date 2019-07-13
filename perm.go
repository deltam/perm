// copyright author MISUMI Masaru(deltam)

// Package perm provides a permutation generator based on group theory.
package perm

import (
	"reflect"
)

// Perm represents current permutation
type Perm struct {
	cur   []int
	slice interface{}
	done  bool

	cycleEnd   []int
	onBigCycle bool
}

// New returns permutation generator
func New(n int) Perm {
	end := make([]int, n)
	for i := 0; i < n; i++ {
		end[i] = n - i - 1
	}
	start := make([]int, n)
	copy(start, end)
	swap(start)
	return Perm{cur: start, cycleEnd: end, onBigCycle: n < 3}
}

// Iter generates slice's permutation generator
func Iter(slice interface{}) Perm {
	rv := reflect.ValueOf(slice)
	len := rv.Len()
	p := New(len)
	p.slice = slice
	return p
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
	end := eqArray(p.cur, p.cycleEnd)
	if end && p.onBigCycle {
		p.done = true
		return
	}
	if doSwap(p.cur) && !end {
		swap(p.cur)
		swapSlice(p.slice)
	} else {
		rot(p.cur)
		rotSlice(p.slice)
	}
	if end && !p.onBigCycle {
		p.onBigCycle = true
		copy(p.cycleEnd, p.cur)
		swap(p.cycleEnd)
	}
}

// Done returns permutation is all
func (p Perm) Done() bool {
	return p.done
}

func rot(p []int) {
	n := len(p)
	f := p[0]
	copy(p[0:n-1], p[1:n])
	p[n-1] = f
}

func swap(p []int) {
	if len(p) < 2 {
		return
	}
	p[1], p[0] = p[0], p[1]
}

func rotSlice(slice interface{}) {
	if slice == nil {
		return
	}
	rv := reflect.ValueOf(slice)
	sw := reflect.Swapper(slice)
	n := rv.Len()
	for i := 1; i < n; i++ {
		sw(i-1, i)
	}
}

func swapSlice(slice interface{}) {
	if slice == nil {
		return
	}
	reflect.Swapper(slice)(0, 1)
}

func doSwap(p []int) bool {
	n := len(p)
	if n < 2 {
		return false
	}
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
	return true
}

func eqArray(s1 []int, s2 []int) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
