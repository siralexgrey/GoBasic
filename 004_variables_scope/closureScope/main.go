package main

import "fmt"

func main() {
	// x := 100
	// fmt.Println(x)
	// someFunc()

	dialog()
}

func someFunc() {
	x := 150
	fmt.Println(x)
	fmt.Println(miniGlobal) // глобальную перемнную можно использовать в любом месте
}

func someFuncTwo() {

	// поряд инициализации и вызова

	var test = "ok!"
	var point = 4.4

	fmt.Println(test)
	fmt.Println(point)
}

var miniGlobal = true // глобальная переменная

func dialog() {
	var x = "Hi"
	var y = "Hello"
	fmt.Println(x)
	{
		fmt.Println(y)
		var z = "How are you?"
		fmt.Println(z)
	}
	// fmt.Println(z) - нельзя из-за вложености
}
