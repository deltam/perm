// Copyright 2019 deltam

package perm

import "reflect"

// Generator is an interface of permutation generator.
type Generator interface {
	Index() []int
	HasNext() bool
	Next() (swapped bool)
}

// permState represents current permutation generator state.
type permState struct {
	cur             []int
	done            bool
	successor       func(*permState) bool
	smallCycleIndex int
	smallCycleSize  int
}

// New returns permutation generator
func New(n int) (Generator, error) {
	start := make([]int, n)
	if err := Init(start); err != nil {
		return nil, err
	}
	return &permState{
		cur:             start,
		successor:       successorSmallCycleShift,
		smallCycleIndex: 0,
		smallCycleSize:  2 * (n - 1),
	}, nil
}

// StartFrom returns permutation generator that start from specified permutation
func StartFrom(index []int) Generator {
	n := len(index)
	cur := make([]int, n)
	copy(cur, index)
	idx := SmallCycleIndex(cur)
	if idx < 0 {
		return &permState{
			cur:       cur,
			successor: successorLargeCycle,
		}
	}
	succ := successorSmallCycleSwap
	if idx%2 == 0 {
		succ = successorSmallCycleShift
	}
	return &permState{
		cur:             cur,
		smallCycleIndex: idx,
		smallCycleSize:  2 * (n - 1),
		successor:       succ,
	}
}

// Index returns current permutation as array
//
// CAUTION: If edit returned array, generator breaks
func (p *permState) Index() []int {
	return p.cur
}

// HasNext returns true if current permutation is not last
func (p *permState) HasNext() bool {
	return !p.done
}

// Next changes current permutation to next
func (p *permState) Next() bool {
	if p.done {
		return false
	}
	return p.successor(p)
}

func successorSmallCycleSwap(p *permState) bool {
	if p.smallCycleIndex >= p.smallCycleSize-1 {
		p.successor = successorLargeCycle
		OpRotate(p.cur)
		return false
	}

	OpSwap(p.cur)
	p.smallCycleIndex++
	p.successor = successorSmallCycleShift
	return true
}

func successorSmallCycleShift(p *permState) bool {
	if p.smallCycleIndex >= p.smallCycleSize-1 {
		p.successor = successorLargeCycle
		OpRotate(p.cur)
		return false
	}

	OpRotate(p.cur)
	p.smallCycleIndex++
	p.successor = successorSmallCycleSwap
	return false
}

func successorLargeCycle(p *permState) bool {
	if IsSwap(p.cur) {
		if IsEnd(p.cur) {
			p.done = true
			return false
		}
		OpSwap(p.cur)
		return true
	}
	OpRotate(p.cur)
	return false
}

type iterator struct {
	g     Generator
	slice interface{}
}

func Bind(g Generator, slice interface{}) Generator {
	return &iterator{g: g, slice: slice}
}

// Iter returns an iterator of slice's all permutation
func Iter(slice interface{}) (Generator, error) {
	rv := reflect.ValueOf(slice)
	len := rv.Len()
	g, err := New(len)
	if err != nil {
		return nil, err
	}
	return Bind(g, slice), nil
}

func (it *iterator) Index() []int {
	return it.g.Index()
}

func (it *iterator) HasNext() bool {
	return it.g.HasNext()
}

func (it *iterator) Next() (swapped bool) {
	if !it.g.HasNext() {
		return false
	}

	swapped = it.g.Next()
	if swapped {
		swapSlice(it.slice)
		return
	}
	shiftSlice(it.slice)
	return
}

func shiftSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	n := rv.Len()
	f := rv.Index(0).Interface()
	reflect.Copy(rv.Slice(0, n-1), rv.Slice(1, n))
	rv.Index(n - 1).Set(reflect.ValueOf(f))
}

func swapSlice(slice interface{}) {
	reflect.Swapper(slice)(0, 1)
}
