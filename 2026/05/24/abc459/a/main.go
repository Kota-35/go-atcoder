// https://atcoder.jp/contests/abc459/tasks/abc459_a
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	word := "HelloWorld" 
	fmt.Println(word[:x-1] + word[x:])
}
