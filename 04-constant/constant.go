package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000 // 숫자 상수는 타ㅓ입이 명확하게 주어지기 전까지는 타입 없음
	const d = 3e20 / n  // 이 연산을 통해 실수 형태의 값을 가지게 되는데 Go언어는 타입이 엄격해서 실수를 정수형 변수에 넣거나 출력 시 타입이 안 맞으면 에러
	fmt.Println(d)
	fmt.Println(int64(d))    // int64는 Go 언어의 정수형 타입 중 하나로 64비트 정수를 의미 <- 그래서 여기에서 현재 실수 형태인 d값을 강제로 64비트 정수형으로 변환! casting
	fmt.Println(math.Sin(n)) // math.Sin은 float64 타입을 기대 -> 데이터타입이 명확히 정해져 있음(타입에 엄격하니까)
	// 그런데 바로 윗 줄이 에러 없이 작동하는 이유? n은 정수인데? -> "타입이 없는 상수" 개념이 있다! 그래서 이 때야 float64타입으로 변환해서 함수에 전달 (컴파일러가 알맞게 변환)
}
