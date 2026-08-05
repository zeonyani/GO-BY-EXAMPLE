package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a) // 대문자는 public, 소문자는 private + int 초기화 안할 시 0 배열

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a)) // 내장함수 len 길이 반환

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
