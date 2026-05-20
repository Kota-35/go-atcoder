// https://atcoder.jp/contests/abc318/tasks/abc318_c
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, d, p int
	fmt.Scan(&n, &d, &p)
	fs := make([]int, n)
	ss := make([]int, n)
	for i := range n {
		fmt.Scan(&fs[i])
	}

	sort.Ints(fs)
	ss[0] = fs[0]
	for i := 0; i < n-1; i++ {
		ss[i+1] = ss[i] + fs[i+1]
	}

	k := (n + d - 1) / d
	ans := p * k
	for i := range k {
		ans = min(ans, ss[n-1-(i*d)]+(p*i))
	}

	fmt.Println(ans)
}
