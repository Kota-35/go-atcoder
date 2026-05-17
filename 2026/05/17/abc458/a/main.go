// https://atcoder.jp/contests/abc458/tasks/abc458_a
package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s)
	fmt.Scan(&n)

	fmt.Println(s[n : len(s)-n])
}
