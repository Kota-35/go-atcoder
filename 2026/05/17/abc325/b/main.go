// https://atcoder.jp/contests/abc325/tasks/abc325_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int

	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	ws := make([]int, n)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		ws[i], _ = strconv.Atoi(inputs[0])
		xs[i], _ = strconv.Atoi(inputs[1])
	}

	answear := 0
	for i := range 24 {

		count := 0
		for j := 0; j < n; j++ {
			// 基準時との時間差
			// 現地の時間は24で割った余り
			td := (i + xs[j]) % 24

			if td >= 9 && td < 18 {
				count += ws[j]
			}
		}

		answear = max(answear, count)
	}

	fmt.Println(answear)
}
