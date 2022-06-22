package main

import (
	"errors"
	"fmt"
)

var a, b, c int

var msg string

func init() {
	msg = "from init()\n"
}

func main() {
	fmt.Println(msg)

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

	/*
		message4, err := faceCheck(15)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf(message4)
	*/

	fmt.Println(findMax(4, 3, 2, 76, 54, 342))

	func() {
		fmt.Println("Anonim func")
	}()

	inc := increment()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())

	message5 := "Go to Golang"
	printMessage1(message5)
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

func faceCheck(age int) (string /*bool,*/, error) {
	if age >= 18 && age < 45 {
		return "Come in" /*true,*/, nil
	} else if age >= 45 && age < 100 {
		return "Go home" /*false,*/, errors.New("age to old")
	}
	return "Go out" /*false,*/, errors.New("age to young")
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}

	return max
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}

func printMessage1(message5 string) {
	message5 += " (from printMessage1())"
	fmt.Println(message5)
}
