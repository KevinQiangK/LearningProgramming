package main

import (
	"fmt"
	"runtime"
)

func basicIfDemo() {
	if runtime.GOOS == "linux" {
		fmt.Println("linux")
	} else {
		fmt.Println(runtime.GOOS)
	}

}

func checkWorkday(day int) {
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("work day", day)
	case 6, 7:
		fmt.Println("weekend day", day)
	default:
		fmt.Println("are you live on earth")
	}

}

func typeSwitchDemo() {
	var x interface{} = 13
	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("the type of x is int", v)
	case string:
		fmt.Println("the type of x is string")
	default:
		fmt.Println("don't support the type")
	}
}

func basicSwitchDemo() {
	checkWorkday(3)
	checkWorkday(7)
	checkWorkday(8)
}

func main() {
	fmt.Println("start learning go if control")

	basicIfDemo()

	basicSwitchDemo()

	typeSwitchDemo()
}
