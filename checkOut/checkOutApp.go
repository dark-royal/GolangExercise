package main

import (
	"fmt"
	"strings"
)

var product []string
var quantity []int
var price []float64
var total float64

var customerName string
var amountPaid float64
var balance float64
var cashierName string
var subTotal float64
var discount float64
var discountedPrice float64
var billTotal float64

func main() {
	collectInput()

}

func collectInput() {
	fmt.Println("What is the customer name: ")
	fmt.Scan(&customerName)

	var moreItem string
	fmt.Println("What did the user buy?")
	var userProduct string
	fmt.Scan(&userProduct)
	product = append(product, userProduct)

	fmt.Println("What is the quantity of what you want to buy?")
	var userQuantity int
	fmt.Scan(&userQuantity)
	quantity = append(quantity, userQuantity)

	fmt.Println("How much per unit?")
	var productPrice float64
	fmt.Scan(&productPrice)
	price = append(price, productPrice)

	fmt.Println("Do you want to add more item?")
	fmt.Scan(&moreItem)
	for {
		if strings.ToLower(moreItem) == "yes" {
			break
		}
	}
	fmt.Println("What is the cashier name")
	fmt.Scan(&cashierName)

	fmt.Println("How much discount will he get?")
	fmt.Scan(&discount)

}

func calculateEachItemTotalPrice() {
	for i := 0; i < len(product); i++ {
		total = append(total, float64(quantity[i])*price[i])

	}

}
