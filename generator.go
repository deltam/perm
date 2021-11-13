// Copyright 2019 deltam

package perm

import "reflect"

// Generator is an interface of permutation generator.
type Generator interface {
	Index() []int
	Done() bool
	Next()
}

// permState represents current permutation generator state.
type permState struct {
	cur             []int
	done            bool
	inSmallCycle    bool
	smallCycleIndex int
}

// New returns permutation generator
func New(n int) (Generator, error) {
	start := make([]int, n)
	if err := Init(start); err != nil {
		return nil, err
	}
	return &permState{
		cur:          start,
		inSmallCycle: true,
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
			cur:          cur,
			inSmallCycle: false,
		}
	}
	return &permState{
		cur:             cur,
		inSmallCycle:    true,
		smallCycleIndex: idx,
	}
}

// Index returns current permutation as array
//
// CAUTION: If edit returned array, generator breaks
func (p *permState) Index() []int {
	return p.cur
}

// HasNext returns true if current permutation is not last
func (p *permState) Done() bool {
	return p.done
}

// Next changes current permutation to next
func (p *permState) Next() {
	// Small Cycle Rule
	if p.inSmallCycle {
		if p.smallCycleIndex < 2*(len(p.cur)-1)-1 {
			if p.smallCycleIndex%2 == 0 {
				OpRotate(p.cur)
			} else {
				OpSwap(p.cur)
			}
			p.smallCycleIndex++
			return
		}
		p.inSmallCycle = false
		OpRotate(p.cur)
		return
	}

	// End Rule
	if !p.done && IsEnd(p.cur) {
		p.done = true
	}
	if p.done {
		return
	}

	// Large Cycle Rule
	if IsSwap(p.cur) {
		OpSwap(p.cur)
		return
	}
	OpRotate(p.cur)
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

func (it *iterator) Done() bool {
	return it.g.Done()
}

func (it *iterator) Next() {
	if it.g.Done() {
		return
	}

	third := it.g.Index()[2]
	it.g.Next()
	if third == it.g.Index()[2] {
		swapSlice(it.slice)
		return
	}
	rotateSlice(it.slice)
}

func rotateSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	n := rv.Len()
	f := rv.Index(0).Interface()
	reflect.Copy(rv.Slice(0, n-1), rv.Slice(1, n))
	rv.Index(n - 1).Set(reflect.ValueOf(f))
}

func swapSlice(slice interface{}) {
	reflect.Swapper(slice)(0, 1)
}
