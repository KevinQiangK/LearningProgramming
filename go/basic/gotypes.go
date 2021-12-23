package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

type myInt int32

type myIntAlias = int32

func valueTypesDemo() {
	var a, b = int(5), uint(6)
	var p uintptr = 0x1234567

	fmt.Println("signed interger a's length is:", unsafe.Sizeof(a))
	fmt.Println("unsigned interger b's length is:", unsafe.Sizeof(b))
	fmt.Println("uintptr's length is", unsafe.Sizeof(p))

	var m int = 5
	var n int32 = 6

	// var a2 myInt = m
	var a2 myInt = myInt(m)
	var b2 myInt = myInt(n)

	var c myIntAlias = n
	fmt.Println("myInt a2:", a2)
	fmt.Println("myInt b2:", b2)
	fmt.Println("myIntAlias c:", c)
}

func encodeRune() {
	var r rune = 0x4e2d
	fmt.Printf("the unicode character is %c\n", r)

	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)
	fmt.Printf("utf-8 representation is 0x%X\n", buf)
}

func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("the unicode character after decode is %c\n", r)
}

func stringTypeDemo() {
	var s = "中国人"
	fmt.Printf("the length of %s = %d\n", s, len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i])
	}
	fmt.Printf("\n")

	fmt.Println("the character count in s is", utf8.RuneCountInString(s))
	for _, c := range s {
		fmt.Printf("0x%x ", c)
	}
	fmt.Println("")

	var ch_s = '\u4e2d'
	fmt.Printf("%c\n", ch_s)
}

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

func stringInternalImplDemo() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("data:0x%x len:%d\n", hdr.Data, hdr.Len)
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	dumpBytesArray((*p)[:])
}

func main() {

	valueTypesDemo()

	stringTypeDemo()

	encodeRune()

	decodeRune()

	stringInternalImplDemo()
}
