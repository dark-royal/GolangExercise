package main

import "fmt"

func main() {
	var count int

	for count < 8 {
		var count1 int

		for count1 < 8 {
			if count%2 == 0 {
				fmt.Print(" *")

			} else {
				fmt.Print("* ")
			}
			count1++
		}
		fmt.Println(" ")
		count++
	}
	fmt.Println()

}
