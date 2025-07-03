// BOJ 11382 - 꼬마 정민

package main

import "fmt"

func main() {
	var a, b, c int64

	fmt.Scan(&a, &b, &c)

	fmt.Println(a + b + c)
}

// 올바른 정수 타입 선택하기

// Go 언어에서는 더 큰 숫자를 다룰 수 있도록 int64라는 타입을 제공합니다.

// int: 약 -21억 ~ 21억 (32비트 시스템 기준) 또는 약 -900경 ~ 900경 (64비트 시스템 기준)

// int64: 약 -900경 ~ 900경 (시스템과 상관없이 항상 보장)
