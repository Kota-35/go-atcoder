// https://atcoder.jp/contests/abc459/tasks/abc459_a
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	hello := "HelloWorld"
	fmt.Printf("%s%s\n", hello[:x-1], hello[x:])
}
