package main

import "fmt"

func main() {
	var largest int = 0
	var count int = 0

	for count < 10 {

		fmt.Println("Enter numbers")
		var number int
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}
		count++
	}
	fmt.Println("Largest number is :", largest)
}
