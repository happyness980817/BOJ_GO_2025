// Problem 2753 - 윤년

package main

import "fmt"

func main() {
	var Year int
	fmt.Scan(&Year)

	if Year%4 == 0 && Year%100 != 0 {
		fmt.Println(1)
	} else if Year%100 == 0 && Year%400 != 0 {
		fmt.Println(0)
	} else if Year%400 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}
