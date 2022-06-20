package main

import "fmt"

var a, b, c int

func main() {
	var message = "Go to Go"
	const message1 string = "ffds"

	fmt.Println(message1)
	fmt.Println(message)
	fmt.Println("I will be Go developer")

	// data type - https://metanit.com/go/tutorial/2.3.php

	/*
		a, b, c := 1, 2, 3
		fmt.Println(a, b, c)
	*/

	a, _, c = c, a, 8
	fmt.Println(a, b, c)

	a, b, c = 6, 7, 8
	print()

	message3 := sayHello("Nikita", 17)
	printMessage(message3)
}

func print() {
	fmt.Println(a, b, c)
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s! U`r age is %d", name, age)
}
