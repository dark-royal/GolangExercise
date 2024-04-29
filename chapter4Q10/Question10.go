package main

import "fmt"

func main() {
	var x int = 11
	var y int = 9

	if x < 10 {
		if y < 10 {
			fmt.Println("*****")
		}
	} else {
		fmt.Println("#####")
		fmt.Println("$$$$$")
	}
}
