// BOJ 2588 - 곱셈

package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	units := num2 % 10
	tens := (num2 / 10) % 10
	hundreds := num2 / 100

	val3 := num1 * units

	val4 := num1 * tens

	val5 := num1 * hundreds

	val6 := num1 * num2

	fmt.Println(val3)
	fmt.Println(val4)
	fmt.Println(val5)
	fmt.Println(val6)
}
