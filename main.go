// Problem 2739 - 구구단

package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", N, i, N*i)
	}
}
