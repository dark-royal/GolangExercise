package main

import "fmt"

func main() {
	var largest int
	var secondLargest int
	var count int

	for count <= 10 {

		fmt.Println("Enter numbers")
		var number int
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}

		if number < largest {
			secondLargest = number

		}
		count++
	}
	fmt.Println("The largest is ", largest)
	fmt.Println("The second largest is", secondLargest)
}
