package main

import "fmt"

// Total := 100 - не допускается вне функции, только через var

func main() {

	var a = 5 // можно не объявлять тип при инициализации
	var b, c string = "b", "c"
	var d bool = true

	fmt.Printf("%T - %v\n", a, a)
	fmt.Printf("%T - %v\n", b, b)
	fmt.Printf("%T - %v\n", c, c)
	fmt.Printf("%T - %v\n", d, d)

	e := 100
	e = 101
	f := "f"
	g := true

	fmt.Printf("%T - %v\n", e, e)
	fmt.Printf("%T - %v\n", f, f)
	fmt.Printf("%T - %v\n", g, g)

	var zeroInt int
	var emptyString string
	var boolFalse bool

	fmt.Printf("%T - %v\n", zeroInt, zeroInt)
	fmt.Printf("%T - %v\n", emptyString, emptyString)
	fmt.Printf("%T - %v\n", boolFalse, boolFalse)
}
