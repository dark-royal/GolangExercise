package main

import "fmt"

func main() {

	var newBalance int = 0
	var count int = 0

	for count < 3 {

		fmt.Println("Enter account number")
		var accountNUmber string
		fmt.Scan(&accountNUmber)

		fmt.Println("Enter beginning balance")
		var beginningBalance int = 0
		fmt.Scanf("%d", &beginningBalance)

		fmt.Println("Enter  total of all items")
		var charges int = 0
		fmt.Scanf("%d", &charges)

		fmt.Println("Enter total of all credit")
		var credits int = 0
		fmt.Scanf("%d", &credits)

		fmt.Println("Enter credit limit")
		var creditLimit int = 0
		fmt.Scanf("%d", &creditLimit)

		newBalance = beginningBalance + charges - creditLimit

		if newBalance > creditLimit {
			fmt.Println("credit limit exceeded")
		}
		count++

		fmt.Println("The new balance after the month is", newBalance)
	}
}
