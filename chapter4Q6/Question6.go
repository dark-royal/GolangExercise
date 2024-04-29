package main

import "fmt"

func main() {
	var count1 int
	var count2 int
	var count3 int
	var count4 int

	fmt.Println("N\t\tN2\t\tN3\t\tN4\t\t")
	for count1 < 5 {

		count2 = count1 * count1
		count3 = count2 * count1
		count4 = count2 * count2
		count1++

		fmt.Printf("%d%8d%8d%8d\n", count1, count2, count3, count4)

	}

}
