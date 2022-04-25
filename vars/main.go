package main

import "fmt"

func main() {
	var uname string = "Gaurav"
	fmt.Println(uname)
	fmt.Printf("Type: %T\n", uname)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type: %T\n", smallVal)

	var smallFloat float64 = 255.4343244534354
	fmt.Println(smallFloat)
	fmt.Printf("Type: %T\n", smallFloat)

	//Go Lexar will decide for data type
	var name = "Gaurav"
	fmt.Println(name)

	users := 1000 // Only allowed in  local scope not in global
	fmt.Println(users)
}
