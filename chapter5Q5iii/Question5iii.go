package main

import "fmt"

func main() {

	for count1 := 0; count1 < 10; count1++ {
		for count := 0; count < count1; count++ {
			for i := 0; i < count; i++ {
				fmt.Print()
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}
