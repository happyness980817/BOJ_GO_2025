// Problem 9498 - 시험 성적

package main

import "fmt"

func main() {
	var score int
	fmt.Scan(&score)

	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
