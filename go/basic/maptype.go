package main

import "fmt"

func mapBasicDemo() {
	// var m map[string]string //panic: assignment to entry in nil map

	var m = map[string]string{}
	m["key"] = "hello"

	m1 := map[int][]string{
		1: {"go", "c++", "python"},
		3: {"learning go", "learning python"},
	}

	fmt.Println(m1)

	type Point struct {
		x float64
		y float64
	}

	m2 := map[Point]string{
		{3.0, 4.0}: "school",
		{5.0, 7.0}: "office",
	}
	fmt.Println(m2)
}

func mapUsingMakeDemo() {
	m1 := make(map[int]string)
	m2 := make(map[int]string, 8)

	m1[1] = "value1"
	m1[10] = "value10"
	m1[1] = "value100"

	fmt.Println(m1, len(m1))
	fmt.Println(m2, len(m2))

	mstr := make(map[string]int)
	mstr["key1"] = 200
	v := mstr["key1"]
	fmt.Println(mstr["key1"], v)
	fmt.Println(mstr["not_exist_key"]) //print 0, the default value

	v, ok := mstr["not_exist_key"]
	if !ok {
		fmt.Println("not exist key", ok)
	}

	fmt.Println(mstr, len(mstr))
	delete(mstr, "key1")
	delete(mstr, "not_exist_key")
	fmt.Println(mstr, len(mstr))
}

func foreachMapDemo() {
	m := map[int]int{
		1: 100,
		2: 200,
		3: 300,
	}
	fmt.Printf("{")
	for k, v := range m {
		fmt.Printf("[%d, %d]", k, v)
	}
	fmt.Printf("}\n")

	//程序逻辑不能依赖便利map的元素次序，是无序的
	for k := range m {
		fmt.Println(k)
	}
}

func passUsingRefDemo(m map[string]int) {
	m["key1"] = 100
	m["key2"] = 200
}

func main() {

	mapBasicDemo()

	mapUsingMakeDemo()

	foreachMapDemo()

	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(m)
	passUsingRefDemo(m)
	fmt.Println(m)

}
