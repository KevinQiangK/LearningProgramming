package main

import (
	"fmt"
	"unsafe"
)

func arrayLengthDemo() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("array length is %d\n", len(arr))
	fmt.Printf("array byte length is %d\n", unsafe.Sizeof(arr))

	var arr2 = [...]int{11, 12, 13}
	fmt.Printf("%T\n", arr2)

	var arr3 = [...]int{
		3: 30,
		9: 100,
	}
	fmt.Printf("%T\n", arr3)

}

func sliceDemo() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Printf("slice length is %d, cap is %d\n", len(numbers), cap(numbers))
	numbers = append(numbers, 6, 7)
	fmt.Printf("slice length is %d, cap is %d\n", len(numbers), cap(numbers))

	sl := make([]byte, 6, 10)
	fmt.Printf("slice length is %d, cap is %d\n", len(sl), cap(sl))

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl2 := arr[3:7:9]
	sl3 := arr[1:4]
	fmt.Printf("slice length is %d, cap is %d\n", len(sl2), cap(sl2))
	fmt.Printf("slice length is %d, cap is %d\n", len(sl3), cap(sl3))

	fmt.Printf("arr[3] = %d\n", arr[3])
	sl2[0] += 10
	fmt.Printf("arr[3] = %d\n", arr[3])
	sl3[2] = 100
	fmt.Printf("arr[3] = %d\n", arr[3])

	u := [...]int{11, 12, 13, 14, 15}
	fmt.Println("array:", u)
	s := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v \n", len(s), cap(s), s)

	s = append(s, 24)
	fmt.Printf("after append 24, array:%v \n", u)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v \n", len(s), cap(s), s)

	s = append(s, 25)
	fmt.Printf("after append 25, array:%v \n", u)
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v \n", len(s), cap(s), s)

	s = append(s, 26)
	fmt.Printf("after append 26, array:%v \n", u)
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v \n", len(s), cap(s), s)

	s[0] = 22
	fmt.Println("after reassign 1st elment of slice, array:", u)
	fmt.Printf("after reassign 1st elment of slice, slice(len=%d, cap=%d): %v \n", len(s), cap(s), s)

}

func sliceDiffDemo() {
	var s1 []int
	var s2 = []int{}

	fmt.Println("s1 声明 value is nil. s1 == nil ? ", s1 == nil)
	fmt.Println("s2 初始化 value is empty. s2 == nil ? ", s2 == nil)

	fmt.Printf("slice(len=%d, cap=%d， addr=%p): %v \n", len(s1), cap(s1), &s1, s1)

	fmt.Printf("slice(len=%d, cap=%d,  addr=%p): %v \n", len(s2), cap(s2), &s2, s2)

}

func main() {
	arrayLengthDemo()
	sliceDemo()
	sliceDiffDemo()
}
