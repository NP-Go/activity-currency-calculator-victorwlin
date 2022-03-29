package main

import (
	"fmt"
	"math"
)

func CurrencyCalculator(oneDollar, fiftyCent, twentyCent, tenCent, fiveCent float64) (float64, float64, float64) {
	//Insert your code here
	totalAmount := oneDollar * 1 + fiftyCent * 0.5 + twentyCent * 0.2 + tenCent * 0.1 + fiveCent * 0.05
	twoDollarNotes := math.Trunc(totalAmount / 2)
	changeAmount := math.Mod(totalAmount, 2)

	//Do not remove this
	fmt.Println("Total:", totalAmount, "Two Dollar Notes:", twoDollarNotes, " Change: ", changeAmount)
	return totalAmount, twoDollarNotes, changeAmount
}

func main() {
	var oneDollarCoin float64
	var fiftyCentCoin float64
	var twentyCentCoin float64
	var tenCentCoin float64
	var fiveCentCoin float64

	fmt.Printf("Enter the number of 1-dollar coins: ")
	fmt.Scanln(&oneDollarCoin)
	fmt.Printf("Enter the number of 50-cent coins: ")
	fmt.Scanln(&fiftyCentCoin)
	fmt.Printf("Enter the number of 20-cent coins: ")
	fmt.Scanln(&twentyCentCoin)
	fmt.Printf("Enter the number of 10-cent coins: ")
	fmt.Scanln(&tenCentCoin)
	fmt.Printf("Enter the number of 5-cent coins: ")
	fmt.Scanln(&fiveCentCoin)
	CurrencyCalculator(oneDollarCoin, fiftyCentCoin, twentyCentCoin, tenCentCoin, fiveCentCoin)
}
