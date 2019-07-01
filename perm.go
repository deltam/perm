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
		delta(p.cur)
		deltaSlice(p.slice)
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
	for i := 0; i < n-2; i++ {
		start[i] = n - i - 2
	}
	start[n-2] = n - 1
	start[n-1] = 0
	return start
}

func (p Perm) isLast() bool {
	n := len(p.cur)
	if p.cur[0] != 0 || p.cur[1] != 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if p.cur[i] != n-i+1 {
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

// 2 4 5 ... n-1 1 0
func delta(p []int) {
	n := len(p)
	f0 := p[0]
	f1 := p[1]
	copy(p[0:n-2], p[2:n])
	p[n-2] = f1
	p[n-1] = f0
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

func deltaSlice(slice interface{}) {
	if slice == nil {
		return
	}
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	n := rv.Len()
	for i := 2; i < n; i++ {
		swap(i-2, i)
	}
	if n%2 == 0 {
		swap(n-2, n-1)
	}
}

func check(p []int) bool {
	n := len(p)
	if p[0] == n-1 {
		return false
	}
	pos := 0
	for i := 0; i < n; i++ {
		if p[i] == n-1 {
			pos = i
			break
		}
	}
	pos++
	if pos == n {
		pos = 1
	}
	if p[pos] != (p[0]-2+n)%(n-1) {
		return false
	}
	for i := 0; i < n; i++ {
		if p[(i+1)%n] != n-i-1 {
			return true
		}
	}
	return false
}
