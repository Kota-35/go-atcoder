// https://atcoder.jp/contests/abc459/tasks/abc459_b
package main

import (
	"fmt"
	"strings"
)


func solve(s string) string {
	if s[0] == 'a' ||  s[0] == 'b' || s[0] == 'c' {
		return "2"
	} else if s[0] == 'd' || s[0] == 'e' || s[0] == 'f' {
		return "3"
	} else if s[0] == 'g' || s[0] == 'h' || s[0] == 'i' {
		return "4"
	}else if s[0] == 'j' || s[0] == 'k' || s[0] == 'l' {
		return "5"
	}else if s[0] == 'm' || s[0] == 'n' || s[0] == 'o' {
		return "6"
	}else if s[0] == 'p' || s[0] == 'q' || s[0] == 'r' || s[0] == 's'{
		return "7"
	}else if s[0] == 't' || s[0] == 'u' || s[0] == 'v' {
		return "8"
	}else if s[0] == 'w' || s[0] == 'x' || s[0] == 'y' || s[0] == 'z' {
		return "9"
	}
	
	return "1"

}
func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range n {
		fmt.Scan(&ss[i])
	}

	cs := make([]string, n)
	for i, s := range ss {
		cs[i] = solve(s)
	}

	fmt.Println(strings.Join(cs, ""))
}
