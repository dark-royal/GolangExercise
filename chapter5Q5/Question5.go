package main

import "fmt"

func main() {

	for count := 10; count >= 1; count-- {
		for i := 0; i < count; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
