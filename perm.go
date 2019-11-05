// Copyright 2019 deltam

// Package perm provides a permutation generator based on group theory.
package perm

import (
	"reflect"
)

// Perm represents current permutation
type Perm struct {
	cur        []int
	slice      interface{}
	done       bool
	largeCycle bool
}

// New returns permutation generator
func New(n int) *Perm {
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
	if n <= 2 {
		p.largeCycle = true
	}
	return &p
}

// Iter returns an iterator of slice's all permutation
func Iter(slice interface{}) *Perm {
	rv := reflect.ValueOf(slice)
	len := rv.Len()
	p := New(len)
	p.slice = slice
	return p
}

// StartFrom returns permutation generator that start from specified permutation
func StartFrom(index []int, slice interface{}) *Perm {
	cur := make([]int, len(index))
	copy(cur, index)
	largeCycle := true
	// Does index belong to Large Cycle?
	if n := len(index); n > 2 && (index[0] == n-1 || index[1] == n-1) {
		for i := 2; i < n+2; i++ {
			j := (i + 1) % n
			if index[j] == n-1 {
				j++
			}
			if index[i%n]-1 != index[j] {
				largeCycle = false
				break
			}
		}
	}
	return &Perm{cur: cur, slice: slice, largeCycle: largeCycle}
}

// Index returns current permutation as array index
func (p *Perm) Index() []int {
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
	if n < 2 || p.largeCycle && isDescOrder(p.cur, n-1) {
		p.done = true
		return
	}
	smallCycleEnd := !p.largeCycle && isDescOrder(p.cur, n)
	successor(p, smallCycleEnd)
	if smallCycleEnd && n > 2 {
		successor(p, false)
		p.largeCycle = true
	}
}

// Done returns permutation is all
func (p *Perm) Done() bool {
	return p.done
}

func successor(p *Perm, forceRot bool) {
	if len(p.cur) < 2 {
		return
	}
	if !forceRot && ruleSwap(p.cur) {
		swap(p.cur)
		if p.slice != nil {
			swapSlice(p.slice)
		}
	} else {
		rot(p.cur)
		if p.slice != nil {
			rotSlice(p.slice)
		}
	}
}

func rot(p []int) {
	n := len(p)
	f := p[0]
	copy(p[0:n-1], p[1:n])
	p[n-1] = f
}

func swap(p []int) {
	p[1], p[0] = p[0], p[1]
}

func rotSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	n := rv.Len()
	f := rv.Index(0).Interface()
	reflect.Copy(rv.Slice(0, n-1), rv.Slice(1, n))
	rv.Index(n - 1).Set(reflect.ValueOf(f))
}

func swapSlice(slice interface{}) {
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
	if p[0] == n-1 {
		return p[2] == p[1]-1 || p[1] == 0 && p[2] == n-2
	}
	if p[n-1] == n-1 {
		return p[0] == p[1]-1 || p[1] == 0 && p[0] == n-2
	}
	for i := 2; i < n-1; i++ {
		if p[i] == n-1 {
			return p[i+1] == p[1]-1 || p[1] == 0 && p[i+1] == n-2
			//return p[i+1] == (p[1]-1+(n-1))%(n-1)
		}
	}
	return false // not reached
}

func isDescOrder(p []int, n int) bool {
	m := n - 1
	for i := 0; i < n; i++ {
		if p[i] != m {
			return false
		}
		m--
	}
	return true
}
