package main

import (
	"fmt"
	"strings"
	"time"
)

var product []string
var quantity []int
var price []float64
var total []float64

var customerName string
var discount1 float64

var amountPaid float64
var balance float64
var cashierName string
var subTotal float64
var discount float64
var discountedPrice float64
var vat float64
var billTotal float64

func main() {
	collectInput()
	calculateEachItemTotalPrice()
	calculateAllTotal()
	calculateDiscountAmount()
	calculateVat()
	getBillTotal()
	printFirstReceiptAfterTheCustomerPaid()
	amountUserPaid()
	getBalance()
	printPaymentReceipt()

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

func calculateAllTotal() {
	for i := 0; i < len(product); i++ {
		subTotal += total[i]
	}
}

func calculateDiscountAmount() float64 {
	discount = (discount / 100) * subTotal
	discount1 = subTotal - discount
	return discount1

}

func calculateVat() float64 {
	vat = (17.50 / 100) * subTotal
	return vat
}

func getBillTotal() float64{
	billTotal = (subTotal - calculateDiscountAmount()) + calculateVat()
	return billTotal
}

func printFirstReceiptAfterTheCustomerPaid() {
	fmt.Printf(`
	SEMICOLON STORES +
	MAIN BRANCH
	LOCATION:312, HERBERT MACAULAY WAY, sABO yABA, LAGOS.
	TEL:08364537281
	Date:%s%n

`, time.Now())
	fmt.Printf("Cashier:%s\n customerName:%s\n", cashierName, customerName)
	fmt.Println("====================================================================")
	fmt.Printf("		%5s%12s%9s%12s\n\n", "ITEMS", "QUANTITY", "PRICE", " TOTAL")
	fmt.Println("---------------------------------------------------------------------")
	for i, prod := range product {
		fmt.Printf("        %s%9d%9.2f%12.2f\n\n", prod, quantity[i], price[i], total[i])
	}
	fmt.Println("---------------------------------------------------------------------")

	fmt.Printf(`
	sub total : %.2f
	
	discount: %.2f

	VAT @ 17.50: %.2f

	bill total : %.2f

	==================================================================================

			THIS IS NOT A RECEIPT KINDLY PAY %.2f

	==================================================================================

`, subTotal, calculateDiscountAmount(), calculateVat(), billTotal, billTotal)
}

func amountUserPaid() float64{
	fmt.Print("How much did the user give to you? ")
	fmt.Scan(&amountPaid)
	return  amountPaid

}

func  getBalance() float64 {
	balance = amountUserPaid() - getBillTotal();
	return balance
}

func printPaymentReceipt()  {
	fmt.Printf(`
	SEMICOLON STORES +
	MAIN BRANCH
	LOCATION:312, HERBERT MACAULAY WAY, sABO yABA, LAGOS.
	TEL:08364537281
	Date:%s%n

`, time.Now())
	fmt.Printf("Cashier:%s\n customerName:%s\n", cashierName, customerName)
	fmt.Println("====================================================================")
	fmt.Printf("		%5s%12s%9s%12s\n\n", "ITEMS", "QUANTITY", "PRICE", " TOTAL")
	fmt.Println("---------------------------------------------------------------------")
	for i, prod := range product {
		fmt.Printf("        %s%9d%9.2f%12.2f\n\n", prod, quantity[i], price[i], total[i])
	}
	fmt.Println("---------------------------------------------------------------------")

	fmt.Printf(`
	sub total : %.2f
	
	discount: %.2f

	VAT @ 17.50: %.2f

	==================================================================================
	bill total : %.2f

	amount paid:%.2f

	balance: %.2f
	==================================================================================

	

	==================================================================================

			Thanks for your patronage

	==================================================================================

	
}













