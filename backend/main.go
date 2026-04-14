package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")

	// var name string = "sd"
	// var age = 18
	// city := "shanghai"
	// var store int
	// fmt.Println(name, age, city, store)

	// const Pi = 3.14
	// const  (
	// 	success = 200
	// 	notFound = 404
	// )

	// const a = 1
	// var b = float32(a)
	// fmt.Println(b)

	const score = 80
	if score := 85 ; score > 80 {
		fmt.Println("Pass")
	} else if score == 80 {
		fmt.Println("Perfect")
	} else {
		fmt.Println("Fail")
	}
}
