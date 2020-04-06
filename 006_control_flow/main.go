package main

import "fmt"

func main() {
	// if / else

	// a, b, c := 10, 20, 30

	// if a > b {
	// 	fmt.Println("a")
	// } else if b < c {
	// 	fmt.Println("b")
	// } else {
	// 	fmt.Println("c")
	// }

	// var text = "admin"

	// switch text {
	// case "admin":
	// 	fmt.Println("Error not admin")
	// case "odmin":
	// 	fmt.Println("Error not admin")
	// case "Admin":
	// 	fmt.Println("Yeah - Admin")
	// 	fallthrough // проверяет остальные кейсы и выполнит дефолт
	// default:
	// 	fmt.Println("Some login")
	// }

	// switch {
	// case text == "Admin" || text == "admin":
	// 	fmt.Println("ok")
	// default:
	// 	fmt.Println("not ok")
	// }

	// for x := 1; x < 100; x++ {
	// 	if x%10 == 0 {
	// 		fmt.Println(x)
	// 		// continue // завершает итерацию
	// 		break // завершит цикл
	// 	}

	// 	fmt.Println(x)
	// }

	// бесконечный цикл
	// x := 1
	// for x > 0 {
	// 	fmt.Println("x > 0")
	// }

	for x := 1; x > 0; x++ {
		if x%10 == 0 {
			fmt.Println("%10", x)
			break
		}
		fmt.Println("x > 0")
	}
}
