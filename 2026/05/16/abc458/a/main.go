// https://atcoder.jp/contests/abc458/tasks/abc458_a
package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &n)

	fmt.Println(s[n : len(s)-n])
}
