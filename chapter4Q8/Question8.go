package main

import "fmt"

func main() {

	fmt.Println("Enter number")
	var number int
	fmt.Scan(&number)

	for number != 1 && number != 2 {

		fmt.Println("Enter number")
		fmt.Scan(&number)

	}
}
