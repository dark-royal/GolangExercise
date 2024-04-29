package main

import "fmt"

func main() {
	var x int = 5
	var y int = 8

	if y == 8 {
		if x == 5 {
			fmt.Println("@@@@@")
		}
	} else {
		fmt.Println("#####")
		fmt.Println("$$$$$")
		fmt.Println("&&&&&")
	}
}
