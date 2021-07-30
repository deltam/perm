package perm

import (
	"errors"
)

func InitStart(p []int) error {
	n := len(p)
	if n < 3 {
		return errors.New("len(p) >= 3")
	}
	for i := 0; i < n; i++ {
		p[i] = n - i
	}
	swap(p)
	return nil
}

func InitEnd(p []int) {
	n := len(p)
	for i := 0; i < n; i++ {
		p[i] = n - i
	}
	rot(p)
	swap(p)
}

func IsStart(p []int) bool {
	n := len(p)
	return p[0] == n-1 && p[1] == n && isReverse(p[2:])
}

func IsEnd(p []int) bool {
	n := len(p)
	return p[0] == n-2 && p[1] == n-1 && isReverse(p[2:n-1])
}

type Move int8

const (
	Halt Move = iota
	Swap
	Rotation
	RevRotation
)

func (m Move) String() string {
	switch m {
	case Halt:
		return "HALT"
	case Swap:
		return "SWAP"
	case Rotation:
		return "ROT"
	case RevRotation:
		return "REV"
	}
	return "NOT FOUND!"
}

func NextMove(p []int) Move {
	if IsEnd(p) {
		return Halt
	}
	if IsInterchange(p) {
		return Rotation
	}
	if IsSwap(p) {
		return Swap
	}
	return Rotation
}

func IsInterchange(p []int) bool {
	return isReverse(p)
}

func IsSwap(p []int) bool {
	n := len(p)
	m := p[1]
	r := p[2]
	if p[0] != n && p[1] != n {
		for i := 2; i < n; i++ {
			if p[i] == n {
				r = p[(i+1)%n]
				break
			}
		}
	}

	return m == r%(n-1)+1
}

func IsInterchangePrev(p []int) bool {
	return isReverse(p[:len(p)-1])
}

func PrevMove(p []int) Move {
	if IsStart(p) {
		return Halt
	}
	if IsInterchangePrev(p) {
		return RevRotation
	}
	if IsSwapPrev(p) {
		return Swap
	}
	return RevRotation
}

func IsSwapPrev(p []int) bool {
	n := len(p)

	var r int
	for i := 0; i < n; i++ {
		if p[i] == n {
			r = p[i%(n-1)+1]
			break
		}
	}

	return p[0] == r%(n-1)+1
}

func DoMove(s []int, m Move) {
	switch m {
	case Rotation:
		rot(s)
	case Swap:
		swap(s)
	case RevRotation:
		rev(s)
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

func rev(p []int) {
	n := len(p)
	f := p[n-1]
	copy(p[1:n], p[0:n-1])
	p[0] = f
}

func isReverse(p []int) bool {
	n := len(p)
	for i := 0; i < n; i++ {
		if p[i] != n-i {
			return false
		}
	}
	return true
}
