package main

import (
	"fmt"
)

func myFunc() {
	fmt.Println("I am myFunc")
}

func main() {
	var test string = "test"

	var test1 = "test1" //inferred

	var test2 string
	test2 = "test2" //inferred

	test3 := "test3" //inferred

	var q, w, e int = 1, 2, 3 //mutiple variables declaration with same type
	var f, s, g = 11, 12, 13  //mutiple variables declaration with different types

	t, k, h := 23, 24, 25 //mutiple variables declaration using :=

	fmt.Println(q, w, e, f, s, g, t, k, h)

	// Grouping variable declarations
	var (
		a int    = 1
		b string = "This is a string"
		c        = 24
	)

	fmt.Println(a, b, c)
	fmt.Println(test)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println("I go, you go. Ta Ta Bye Bye...!")

	myFunc()
	myAnotherFunc()

}

func myAnotherFunc() {
	fmt.Println("I am myAnotherFunc")
}
