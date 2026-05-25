// https://atcoder.jp/contests/abc458/tasks/abc458_c
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	ans := 0
	for i := range n {
		if s[i] == 'C' {
			ans += min(i, n-i-1) + 1
		}
	}

	fmt.Println(ans)
}
