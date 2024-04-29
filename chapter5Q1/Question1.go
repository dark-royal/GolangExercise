package main

import "fmt"

func main() {

	fmt.Print("Enter a number: ")
	var number int
	fmt.Scan(&number)

	var minimum int = number
	var maximum int
	var sum int
	var count int

	for count < 10 {

		fmt.Print("Enter a number: ")
		var number int
		fmt.Scan(&number)

		if number < minimum {
			minimum = number
		}
		if number > maximum {
			maximum = number
		}

		sum = minimum + maximum
		count++
	}
	fmt.Println("The maximum number is:", maximum)
	fmt.Println("The minimum number is:", minimum)
	fmt.Println("The sum is :", sum)
}
