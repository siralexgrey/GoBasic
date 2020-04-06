package main

import "fmt"

// const applicationName = "App Const"

const (
	_ = iota     // 0
	B = iota     // 1
	C = iota     // 2
	D = iota + 5 // 3 + 5 = 8
)

func main() {

	var temp string = "simple string"
	fmt.Println(temp)
	// fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)

	// const number = 100
	// // fmt.Println(applicationName)
	// fmt.Println(number)

	// var numberTwo int = 10
	// result := number + numberTwo
	// fmt.Println(result)

	x, _, _ := TypesGo()

	// x := TypesGo() - так нельзя, так как функция возвращает 3 значения

	fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
}

func TypesGo() (int, bool, string) {
	return 100, false, "stringType"
}
