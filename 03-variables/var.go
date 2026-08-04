package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // 정수 초기화 안하면 0 기본값, 하나 또는 여러 개 선언 가능
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short" // := 문법: 변수를 선언하면서 초기화 var f string = "short"
	fmt.Println(f)
}
