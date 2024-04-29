package main

import "fmt"

func main() {

	var reversedNumber int
	var originalNumber int
	var calculate int

	fmt.Print("Enter your number: ")
	var number int
	fmt.Scan(&number)

	originalNumber = number

	for number > 0 {
		calculate = number % 10
		number = number / 10
		reversedNumber = (reversedNumber * 10) + calculate

	}

	if reversedNumber == originalNumber {
		fmt.Println("It is a palindrome")

	} else {
		fmt.Println("It is not a palindrome")
	}
}
