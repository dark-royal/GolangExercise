package main

import (
	"fmt"
	"strings"
	"time"
)

type Product struct {
	ProductName string
	quantity    float64
	price       float64
	total       float64
}

var (
	customerName string
	discount1    float64

	amountPaid  float64
	balance     float64
	cashierName string
	subTotal    float64
	discount    float64
	vat         float64
	billTotal   float64
)

var products []Product

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
	_, err2 := fmt.Scan(&customerName)
	if err2 != nil {
		return
	}

	var moreItem string
	fmt.Println("What did the user buy?")
	var userProduct string
	_, err := fmt.Scan(&userProduct)
	if err != nil {
		return
	}

	fmt.Println("What is the quantity of what you want to buy?")
	var userQuantity int
	_, err = fmt.Scan(&userQuantity)
	if err != nil {
		return
	}

	fmt.Println("How much per unit?")
	var productPrice float64
	_, err = fmt.Scan(&productPrice)
	if err != nil {
		return
	}

	product := Product{ProductName: userProduct, quantity: float64(userQuantity), price: productPrice}
	products = append(products, product)

	fmt.Println("Do you want to add more item?")
	_, err2 = fmt.Scan(&moreItem)
	if err2 != nil {
		return
	}
	for {
		if strings.ToLower(moreItem) == "yes" {
			break
		}
	}
	fmt.Println("What is the cashier name")
	_, err = fmt.Scan(&cashierName)
	if err != nil {
		return
	}

	fmt.Println("How much discount will he get?")
	_, err = fmt.Scan(&discount)
	if err != nil {
		fmt.Println("invalid")
		return
	}

}

func calculateEachItemTotalPrice() float64 {
	var total float64
	for _, product := range products {
		total += product.price * product.quantity

	}
	return total

}

func calculateAllTotal() float64 {
	for _, product := range products {
		subTotal += product.total
	}
	return subTotal
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

func getBillTotal() float64 {
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
	for _, prod := range products {
		fmt.Printf("        %s%9d%9.2f%12.2f\n\n", prod.ProductName, prod.quantity, prod.price, prod.total)
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

`, calculateAllTotal(), calculateDiscountAmount(), calculateVat(), getBillTotal(), getBillTotal())
}

func amountUserPaid() float64 {
	fmt.Print("How much did the user give to you? ")
	fmt.Scan(&amountPaid)
	return amountPaid

}

func getBalance() float64 {
	balance = amountUserPaid() - getBillTotal()
	return balance
}

func printPaymentReceipt() {
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
	for _, prod := range products {
		fmt.Printf("        %s%9d%9.2f%12.2f\n\n", prod.ProductName, prod.quantity, prod.price, prod.total)
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


`, calculateAllTotal(), calculateDiscountAmount(), calculateVat(), getBillTotal(), amountUserPaid(), getBalance())
}
