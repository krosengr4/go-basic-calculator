package main

import "fmt"

func main() {
	ifContinue := true

	for ifContinue {
		fmt.Println("1 - Addition(+)\n2 - Subtraction(-)\n3 - Multiplication(*)\n4 - Division(/)\n0 - Quit")
		fmt.Println("Enter the operator you would like to use:")
		var userChoice int
		fmt.Scanln(&userChoice)

		fmt.Println("Enter the first number:")
		var num1 float64
		fmt.Scanln(&num1)

		fmt.Println("Enter the second number:")
		var num2 float64
		fmt.Scanln(&num2)

		switch userChoice {
		case 1:
			add(num1, num2)
		case 2:
			subtract(num1, num2)
		case 3:
			multiply(num1, num2)
		case 4:
			divide(num1, num2)
		case 0:
			ifContinue = false
		}

	}

	fmt.Println("Goodbye!!!")
}

func add(num1, num2 float64) {
	result := num1 + num2
	printResult(result)
}

func subtract(num1, num2 float64) {
	result := num1 - num2
	printResult(result)
}

func multiply(num1, num2 float64) {
	result := num1 * num2
	printResult(result)
}

func divide(num1, num2 float64) {
	if num2 == 0 {
		err := fmt.Errorf("cannot divide by zero")
		fmt.Println(err)
	} else {
		result := num1 / num2
		printResult(result)
	}
}

func printResult(result float64) {
	fmt.Println("Result:", result)
}
