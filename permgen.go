// permgen
// permutation generator
// https://www.gregegan.net/SCIENCE/Superpermutations/Superpermutations.html#LOCAL
// author deltam
package permgen

import "fmt"

type Perm struct {
	cur  []int
	flag bool
}

func NewPerm(n int) Perm {
	inter := make([]int, n)
	for i := 0; i < n; i++ {
		inter[(i+1)%n] = n - i
	}
	g := Perm{cur: inter}
	delta(g.cur)
	return g
}

func (p *Perm) Next() {
	if p.Done() {
		return
	}
	if !p.flag && p.afterLast() {
		p.flag = true
	}
	if check(p.cur) {
		delta(p.cur)
	} else {
		sigma(p.cur)
	}
}

func (p Perm) Done() bool {
	return p.flag && p.afterLast()
}

func (p Perm) afterLast() bool {
	n := len(p.cur)
	for i := 0; i < n; i++ {
		if p.cur[i] != n-i {
			return false
		}
	}
	return true
}

func (p Perm) String() string {
	return fmt.Sprint(p.cur)
}

func fact(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

// 2 3 4 ... n-1 n 1
func sigma(p []int) {
	n := len(p)
	f := p[0]
	copy(p[0:n-1], p[1:n])
	p[n-1] = f
}

// 3 4 5 ... n-1 n 2 1
func delta(p []int) {
	n := len(p)
	f0 := p[0]
	f1 := p[1]
	copy(p[0:n-2], p[2:n])
	p[n-2] = f1
	p[n-1] = f0
}

func check(p []int) bool {
	n := len(p)
	// 1.
	if p[0] == n {
		return false
	}
	pos := 0
	for i := 0; i < n; i++ {
		if p[i] == n {
			pos = i
			break
		}
	}
	pos++
	if pos == n {
		pos = 1
	}
	// 2.
	if p[pos] != 1+(p[0]-2+n-1)%(n-1) {
		return false
	}
	// 3.
	for i := 0; i < n; i++ {
		if p[(i+1)%n] != n-i {
			return true
		}
	}
	return false
}
