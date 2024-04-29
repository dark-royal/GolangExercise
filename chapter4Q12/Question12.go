package main

import "fmt"

func main() {

	fmt.Println("Enter a number")
	var number int
	fmt.Scan(&number)

	for count1 := 0; count1 < number; count1++ {
		for count2 := 0; count2 < count1; count2++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
