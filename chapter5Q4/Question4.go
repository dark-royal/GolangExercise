package main

import "fmt"

func main() {

	for count := 0; count < 10; count++ {
		for i := 0; i < count; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
