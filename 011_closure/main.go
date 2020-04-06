package main

import "fmt"

func main() {

	// result := func(x, y int) int {
	// 	return x + y
	// }(10, 20)

	// fmt.Println(result)

	// fmt.Println(sum(1, 3))

	closure()

	oneMore := func() int {
		return closure() + 5
	}

	fmt.Println(oneMore())
}

func closure() int {
	t := 100

	result := func() {
		fmt.Println(t)
	}

	// z := t + 10

	second := func() {
		fmt.Println(t)
	}

	result()
	second()

	return t
}

// func closure(s string) {

// 	text := "sentece " + " - " + s

// 	fmt.Println(text)

// 	result := func() string {
// 		fmt.Println(text + " some another text")
// 		return text + " some another text"
// 	}

// 	s = "_abc"
// 	sq := func() {
// 		fmt.Println("ahahah" + s)
// 	}

// 	sq()
// 	fmt.Println(result())
// }

// func sum(a, b int) int {

// 	func(s string) {
// 		fmt.Println(s)
// 	}("some text")

// 	return a + b
// }
