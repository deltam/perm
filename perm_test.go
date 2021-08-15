package perm_test

import (
	"fmt"
	"log"

	"github.com/deltam/perm"
)

func Example() {
	p := make([]int, 3)
	if err := perm.Init(p); err != nil {
		log.Fatal(err)
	}

	for !perm.IsEnd(p) {
		fmt.Println(p)
		perm.Advance(p)
	}
	fmt.Println(p)
	// Output:
	// [1 2 0]
	// [2 0 1]
	// [0 2 1]
	// [2 1 0]
	// [1 0 2]
	// [0 1 2]
}

func ExampleNew() {
	g, err := perm.New(3)
	if err != nil {
		log.Fatal(err)
	}

	for g.HasNext() {
		fmt.Println(g.Index())
		g.Next()
	}
	// Output:
	// [1 2 0]
	// [2 0 1]
	// [0 2 1]
	// [2 1 0]
	// [1 0 2]
	// [0 1 2]
}

func ExampleIter() {
	ss := []rune("abc")
	g, err := perm.Iter(ss)
	if err != nil {
		log.Fatal(err)
	}

	for g.HasNext() {
		fmt.Println(string(ss))
		g.Next()
	}
	// Output:
	// abc
	// bca
	// cba
	// bac
	// acb
	// cab
}
