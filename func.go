package perm

import (
	"errors"
)

func Init(p []int) error {
	n := len(p)
	if n < 3 {
		return errors.New("len(p) >= 3")
	}
	for i := 0; i < n; i++ {
		p[i] = n - i - 1
	}
	swap(p)
	return nil
}

func IsEnd(p []int) bool {
	n := len(p)
	return p[0] == n-3 && p[1] == n-2 && isReverse(p[2:n-1])
}

func IsInterchange(p []int) bool {
	return isReverse(p)
}

func IsSwap(p []int) bool {
	n := len(p)
	m := p[1]
	r := p[2]

	if p[0] != n-1 && p[1] != n-1 {
		for i := 2; i < n; i++ {
			if p[i] == n-1 {
				r = p[(i+1)%n]
				break
			}
		}
	}
	return m == (r+1)%(n-1)
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

type Move int8

const (
	Halt Move = iota
	Rotation
	Swap
)

func (m Move) String() string {
	switch m {
	case Halt:
		return "HALT"
	case Rotation:
		return "ROT"
	case Swap:
		return "SWAP"
	}
	return "NOT FOUND!"
}

func DoMove(s []int, m Move) {
	switch m {
	case Rotation:
		rot(s)
	case Swap:
		swap(s)
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
