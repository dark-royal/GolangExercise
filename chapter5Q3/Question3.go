package main

import (
	"fmt"
	"math"
)

func main() {

	principal := 1000.0

	for rate := 0.05; rate <= 0.10; rate += 0.01 {

		fmt.Printf("Interest Rate: %.0f%%\n", rate*100)
		fmt.Println("Year \t\tAmount on deposit")

		for year := 1; year <= 10; year++ {
			amount := principal * math.Pow(1.0+rate, float64(year))
			fmt.Printf("%4d %20.2f\n", year, amount)
		}
		fmt.Println()
	}

}
