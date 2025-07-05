// Problem 14681 - 사분면 고르기

package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x > 0 && y > 0 {
		fmt.Println(1)
	} else if x > 0 && y < 0 {
		fmt.Println(4)
	} else if x < 0 && y > 0 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}
