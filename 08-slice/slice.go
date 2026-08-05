package main

import "fmt"

func main() {
	s := make([]string, 3) // make를 사용하면 배열이 아닌 slice가 만들어짐. 원소의 갯수가 아니라 "포함된 원소들로만" 작성
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))
	s = append(s, "d") // slice에서 적용되는 append 함수
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)
	c := make([]string, len(s))
	copy(c, s) // s를 c로 복사 (뒤를 앞으로)
	fmt.Println("cpy: ", c)
	l := s[2:5]
	fmt.Println("sli: ", l)
	fmt.Println("sli hyang: ", s[2:5])
	l = s[:5]
	fmt.Println("sli2: ", l)
	l = s[2:]
	fmt.Println("sli3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
