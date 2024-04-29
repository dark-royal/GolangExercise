package main

import "fmt"

func main() {
	var count int = 0
	var milesPerGallon float64 = 0
	var totalMilePerGallon float64 = 0

	for count <= 3 {
		fmt.Print("Enter Miles : ")
		var miles float64
		fmt.Scanf("%f", &miles)

		fmt.Println("Enter gallon ")
		var gallon float64
		fmt.Scanf("%f", &gallon)

		milesPerGallon = miles / gallon
		fmt.Println("The miles per gallon is : ", milesPerGallon)

		totalMilePerGallon += milesPerGallon
		count++
	}
	fmt.Print("The total miles per gallon is : ", totalMilePerGallon)

}
