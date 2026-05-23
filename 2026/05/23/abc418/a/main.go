// https://atcoder.jp/contests/abc418/tasks/abc418_a
package main

import  "fmt"


func isTea(s string) bool {
	return s == "tea"
}

func processTea(s string) {
	if isTea(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	if n-3 < 0 {
	processTea(s)	
	} else {
		processTea(s[n-3:])
	}
}
