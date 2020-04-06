package main

import "fmt"

func main() {

	var simpleMap = map[int]string{}
	fmt.Println(simpleMap)
	simpleMap[198] = "some words"
	fmt.Println(simpleMap)

	var anotherMap = make(map[string]int)
	anotherMap["one"] = 198
	anotherMap["two"] = 1981
	anotherMap["three"] = 1982
	fmt.Println(anotherMap)

	delete(anotherMap, "three")
	fmt.Println(anotherMap["two"])

	slice := []int{1, 2, 4}
	for _, v := range slice {
		fmt.Println(v)
		if v == 2 {
			fmt.Println("ok")
		}
	}

	if _, ok := anotherMap["one"]; ok {
		// ok == true
		fmt.Println(anotherMap["one"])
	}
}
