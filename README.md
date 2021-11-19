<div align="center">
<img src="https://raw.githubusercontent.com/deltam/perm/master/img/perm5_ribbon.png" width="100%" alt="n=5 permutation order graph as TAOCP 7.2.1.2 Fig.22 style">
</div>

# perm

[![GoDev]( https://pkg.go.dev/badge/github.com/deltam/perm)](https://pkg.go.dev/github.com/deltam/perm)

Permutation generator based on group theory[^1].

- Fast
- Stateless
    - generate next permutation from current permutation only.
- Permutation order is **NOT** lexical.

## Installation

```
go get github.com/deltam/perm
```

## Usage

Index only:

```go
func main() {
	p := make([]int, 3)
	if err := perm.Init(p); err != nil {
		log.Fatal(err)
	}

	for !perm.IsEnd(p) {
		fmt.Println(p)
		perm.Advance(p)
	}
	fmt.Println(p)
}
// Output:
// [1 2 0]
// [2 0 1]
// [0 2 1]
// [2 1 0]
// [0 1 2]
// [1 0 2]
```

Permutation of slice:

```go
func main() {
	words := []rune("ABC")
	g, err := perm.Iter(ss)
	if err != nil {
		log.Fatal(err)
	}

	for ; !p.Done(); p.Next() {
		fmt.Printf("%v\t%s\n", p.Index(), string(words))
	}
}
// Output:
// [1 2 0]	ABC
// [2 0 1]	BCA
// [0 2 1]	CBA
// [2 1 0]	BAC
// [0 1 2]	CAB
// [1 0 2]	ACB
```

## Performance

This algorithm is 40~50% faster than naive recursive algorithm.

<table class='benchstat '>
<tbody>
<tr><th><th>time/op
<tr><td>Primitive-8<td>33.3µs ± 2%
<tr><td>Generator-8<td>39.1µs ± 2%
<tr><td>PermRecursive-8<td>66.9µs ± 1%
<tr><td>&nbsp;
</tbody>
</table>


```go
func recPerm(p []int, n int, ignore []bool) {
	if n == 0 {
		return
	}
	for i := 0; i < len(p); i++ {
		if !ignore[i] {
			p[n-1] = i
			ignore[i] = true
			recPerm(p, n-1, ignore)
			ignore[i] = false
		}
	}
}
```

## Background

See [paper](https://arxiv.org/abs/1307.2549)[^1] for details.

I also explained details on [my blog post](https://deltam.blogspot.com/2019/12/permutationgenerator.html)(Japanese only). [^3]

[Cayley graph](https://en.wikipedia.org/wiki/Cayley_graph) on permutation group generated by _sigma_ and _tau_ has a size two disjoint cycle cover and [Hamilton path](https://en.wikipedia.org/wiki/Hamiltonian_path).

    sigma(rotation): (1, 2, ..., n) -> (2, 3, ..., n, 1)
    tau(swap):       (1, 2, ..., n) -> (2, 1, ..., n)

<div align="center">
<img src="https://raw.githubusercontent.com/deltam/perm/master/img/cycle_cover4.png" width="80%" alt="cycle cover by n=4">
</div>

Below is the Hamilton path created by split and join the two cycles.

<div align="center">
<img src="https://raw.githubusercontent.com/deltam/perm/master/img/hamilton_path4.png" width="80%" alt="hamilton path by n=4">
</div>

By observing above Hamilton path, you can discover local rule for generating that.

See `perm.ruleSwap()` or this text[^2][^3] for details.

## License

MIT

Copyright 2019 deltam


[^1]: [Hamiltonicity of the Cayley Digraph on the Symmetric Group Generated by σ = (1 2 ... n) and τ = (1 2)](https://arxiv.org/abs/1307.2549)

[^2]: [The Williams Construction: Superpermutations — Greg Egan](https://www.gregegan.net/SCIENCE/Superpermutations/Superpermutations.html#WILLIAMS)

[^3]: [2つの操作のみで全順列を列挙する：対称群のグラフ上のハミルトン路にもとづく順列生成の紹介と実装 : サルノオボエガキ](https://deltam.blogspot.com/2019/12/permutationgenerator.html)
