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
}

// New returns permutation generator
func New(n int) Perm {
	start := make([]int, n)
	for i := 0; i < n; i++ {
		start[i] = n - i - 1
	}
	if n > 2 {
		swap(start)
	}
	p := Perm{cur: start}
	if n < 2 {
		p.done = true
	}
	return p
}

// Iter returns an iterator of slice's all permutation
func Iter(slice interface{}) *Perm {
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
	n := len(p.cur)
	if n < 2 || isDescOrder(p.cur, n-1) {
		p.done = true
		return
	}
	r := isDescOrder(p.cur, n)
	successor(p, r)
	if r && n > 2 {
		successor(p, false)
	}
}

// Done returns permutation is all
func (p Perm) Done() bool {
	return p.done
}

func successor(p *Perm, forceRot bool) {
	if !forceRot && ruleSwap(p.cur) {
		swap(p.cur)
		swapSlice(p.slice)
	} else {
		rot(p.cur)
		rotSlice(p.slice)
	}
}

func rot(p []int) {
	n := len(p)
	if n < 2 {
		return
	}
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
	n := rv.Len()
	if n < 2 {
		return
	}
	f := rv.Index(0).Interface()
	reflect.Copy(rv.Slice(0, n-1), rv.Slice(1, n))
	rv.Index(n - 1).Set(reflect.ValueOf(f))
}

func swapSlice(slice interface{}) {
	if slice == nil {
		return
	}
	reflect.Swapper(slice)(0, 1)
}

func ruleSwap(p []int) bool {
	n := len(p)
	if n <= 2 {
		return false
	}
	if p[1] == n-1 {
		return false
	}
	pos := 2
	for i := 2; i < n; i++ {
		if p[i] == n-1 {
			pos = (i + 1) % n
			break
		}
	}
	if p[pos] != (p[1]+n-2)%(n-1) {
		return false
	}
	return true
}

func isDescOrder(p []int, n int) bool {
	for i := 0; i < n; i++ {
		if p[i] != n-i-1 {
			return false
		}
	}
	return true
}
