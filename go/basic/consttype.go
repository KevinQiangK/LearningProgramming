package main

import "fmt"

type MyInt int32

func constConvertDemo() {
	const n MyInt = 32

	const m int32 = int32(n) + 32

	var a int32 = 5

	fmt.Println(a + int32(n))
}

func constUnTypedDemo() {
	const n = 32

	var a MyInt = 5
	fmt.Println(a + n)
}

func main() {
	constConvertDemo()
	constUnTypedDemo()
}
