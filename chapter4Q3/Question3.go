package main

import "fmt"

func main() {
	var merchandize int = 200.0
	var rate float64 = 0.09
	var count int = 0
	var salesPersonEarning float64 = 0.0

	for count < 2 {

		fmt.Println("Enter item sold")
		var item string
		fmt.Scan(&item)

		fmt.Println("Enter amount")
		var amount float64
		fmt.Scanf("%d", amount)

		salesPersonEarning = amount + float64(merchandize)*rate

		fmt.Println("The sales person earning is ", salesPersonEarning)
		count++
	}
}
