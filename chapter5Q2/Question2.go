package main

import "fmt"

func main() {
	var sum int

	for number := 0; number < 100; number++ {
		if number%3 == 0 {
			sum += number
		}
	}

	fmt.Println("The sum is", sum)

}
