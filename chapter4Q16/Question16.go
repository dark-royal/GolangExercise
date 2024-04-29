package main

import "fmt"

func main() {
	var count int

	for count <= 1 {

		fmt.Println("Enter number")
		var number int
		fmt.Scan(&number)

		if number == number {
			fmt.Println("0")
		} else if number > number {
			fmt.Println("1")

		} else {
			fmt.Println("-1")
		}
		count++
	}

}
