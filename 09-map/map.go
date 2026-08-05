package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7 // 값은 이런 식으로 할당
	m["k2"] = 13

	fmt.Println("map: ", m) // map[키:값] 꼴

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))

	delete(m, "k2") // m이란 map안에 k2라는 이름을 가진 쌍들을 지움
	fmt.Println("map: ", m)

	_, prs := m["k2"]         // 키 존재 여부 확인
	fmt.Println("prs: ", prs) // false 출력
	/*
		맵에 존재하지 않는 키를 조회하면 에러 대신 해당 값 타입의 제로 값 반환
		-> 이 떄 0이라는 게 없어서 0반환 인지, 조회한 값이 0인 건지 파악 곤란
		-> Go 언어에서는 true/false도 함께 반환

		>>이해를 위한 설명 정리<<
		a, b := m["k2"] -----> a에는 맵에 저장된 실제 값(int형 제로 0)
		                -----> b에는 키의 존재 여부 (bool형 false)
		_ 블랭크 식별자 => 사용하지 않는 반환 값을 버리기 위함(남아있으면 컴파일 에러 나서)
	*/
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}
