// https://atcoder.jp/contests/abc318/tasks/abc318_a
package main

import "fmt"

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)

	ans := 0
	for m <= n {
		ans += 1
		m += p
	}

	fmt.Println(ans)
}
