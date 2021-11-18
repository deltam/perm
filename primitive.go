// Copyright 2019 deltam

// Package perm provides a permutation generator based on group theory.
//
// The generated permutations belong to either Small Cycle or Large Cycle.
// The generated permutations starts at small cycle and ends at large cycle.
//
// Small cycle is start from permutation initialized by Init(), end at reverse order(n-1,n-2,...,1,0).
// If applied swap to small cycle end, move to start.
// If applied rotate to small cycle end, move to permutation contained of large cycle.
//
// This algorithm is based on following paper[1] and its commentary article[2] below.
//
// [^1] Hamiltonicity of the Cayley Digraph on the Symmetric Group Generated by σ = (1 2 ... n) and τ = (1 2) https://arxiv.org/abs/1307.2549
//
// [^2] 2つの操作のみで全順列を列挙する：対称群のグラフ上のハミルトン路にもとづく順列生成の紹介と実装 : サルノオボエガキ https://deltam.blogspot.com/2019/12/permutationgenerator.html
package perm

import (
	"errors"
)

var ErrPermutationLengthIsTooShort = errors.New("Permutation length is too short: MUST be len(p) >= 3")

// Init initializes p to a starting permutation consisting of [0,len(p)-1].
// Initialized sequence is n-2,n-1,n-3,...,2,1,0[1].
// If len(p)<3, returns error.
func Init(p []int) error {
	n := len(p)
	if n < 3 {
		return ErrPermutationLengthIsTooShort
	}
	for i := 0; i < n; i++ {
		p[i] = n - i - 1
	}
	OpSwap(p)
	return nil
}

// Advance advances p to next permutation.
// If returns true, the swap operator has been applied.
// Otherwise it is the rotate operator.
func Advance(p []int) (swapped bool) {
	if IsSwap(p) && !IsSmallCycleEnd(p) {
		OpSwap(p)
		return true
	}
	OpRotate(p)
	return false
}

// IsEnd returns true if p is permutation end.
func IsEnd(p []int) bool {
	n := len(p)
	return p[0] == n-3 && p[1] == n-2 && p[n-1] == n-1 && isReverse(p[2:n-1])
}

// IsSwap returns true if next operator is swap.
func IsSwap(p []int) bool {
	n := len(p)
	r := p[0]
	m := p[1]
	if r == n-1 {
		r = p[2]
	} else if m == n-1 {
		return false
	} else {
		for i := 2; i < n-1; i++ {
			if p[i] == n-1 {
				r = p[i+1]
				break
			}
		}
	}
	return m == (r+1)%(n-1)
}

// IsSmallCycleEnd returns true if p is small cycle end.
func IsSmallCycleEnd(p []int) bool {
	return isReverse(p)
}

func isReverse(p []int) bool {
	n := len(p)
	for i := 0; i < n; i++ {
		if p[i] != n-i-1 {
			return false
		}
	}
	return true
}

// OpRotate rotates permutation p.
//  (n-1,n-2,...,1,0) -> (n-2,...,1,0,n-1)
func OpRotate(p []int) {
	n := len(p)
	f := p[0]
	copy(p[0:n-1], p[1:n])
	p[n-1] = f
}

// OpSwap swaps first and second of permutation.
//  (n-1,n-2,...,1,0) -> (n-2,n-1,...,1,0)
func OpSwap(p []int) {
	p[1], p[0] = p[0], p[1]
}
