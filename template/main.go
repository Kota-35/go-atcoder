// https://atcoder.jp/contests/{{{ContestID}}}/tasks/{{{TaskID}}}
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader
var writer *bufio.Writer

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func solve() {

}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	_ = n

	solve()
}
