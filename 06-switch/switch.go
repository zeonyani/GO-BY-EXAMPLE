package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write", i, "  as  ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("오전")
	default:
		fmt.Println("오후")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("BOOL")
		case int:
			fmt.Println("INT")
		default:
			fmt.Println("타입 모름\n", t) // 무조건 한 번 선언하면 꼭 한 번은 호출해야함. time.Now선언과 다르게 switch문에서 선언한 t 스코프
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
