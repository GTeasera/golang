package main

import (
	"errors"
	"fmt"
)

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

func faceCheck(age int) (string, bool, error) {
	if age >= 18 && age < 45 {
		return "Come in", true, nil
	} else if age >= 45 && age < 100 {
		return "Go home", false, errors.New("age to old")
	}
	return "Go out", false, errors.New("age to young")
}
