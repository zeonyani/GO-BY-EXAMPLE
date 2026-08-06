package main

import "fmt"

func main() {
	nums := []int{2, 3, 4} // make로만 슬라이스를 만드는 건 아님
	/*
		slice 선언 정리
		1. nums := []int{2, 3, 4} : 이미 들어갈 초기값을 알고 있는 경우
		2. nums := make([]int, 3, 5) : 길이 3, 용량 5인 슬라이스 생성
			길이(len): 현재 슬라이스에 실제로 들어 있는 데이터 개수
			용량(cap): 메모리상에서 추가 공간 재할당 없이 "최대"담을 수 있는 공간 총 크기
		3. var nums []int : 값은 나중에 넣을 예정이고 당장은 비워두려고 할 때 (값은 nil)
			nil: 아무것도 가리키고 있지 않다는 null의 개념과는 동일하나,
				Go언어에서 여러 타입(슬라이스, 포인터, 맵, 채널, 인터페이스 등)이 가질 수 았는 제로 값
				var nums []int 선언 후 len(nums) 해도 0으로 안전하게 동작
	*/
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums { // 타입 추론을 알아서 함 nums 내에서 키-밸류 타입 보고
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
