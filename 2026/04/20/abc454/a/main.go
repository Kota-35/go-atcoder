// https://atcoder.jp/contests/abc454/tasks/abc454_a
package main

import "fmt"

func solve(l int, r int) int {
	return r - l + 1
}

func main() {
	var l, r int
	fmt.Scanf("%d %d", &l, &r)

	fmt.Print(solve(l, r))
}
