package main

import (
	"fmt"

	"github.com/siralexgrey/GoBasic/002_package/nested/hello"
	"github.com/siralexgrey/GoBasic/002_package/nested/say"
)

func main() {
	fmt.Println("Start programm...")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())
}
