package main

import "fmt"

func main() {
	var taxRate float64
	var count int
	var totalTax float64

	for count < 3 {

		fmt.Println("Enter name")
		var name string
		fmt.Scan(&name)

		fmt.Println("Enter earning")
		var earning float64
		fmt.Scanf("%f", &earning)

		if earning <= 30000 {
			taxRate = 0.15
			totalTax = earning * taxRate

		} else {
			taxRate = 0.20
			totalTax = earning * taxRate
		}

		fmt.Println("The name is", name)
		fmt.Println("The earning is", earning)
		fmt.Println("The total tax is", totalTax)
		count++
	}
}
