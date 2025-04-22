package main

import "fmt"

func main() {
	const USD_EUR float64 = 0.87
	const USD_RUB float64 = 85.53

	var EUR_RUB = USD_RUB / USD_EUR
	fmt.Print(EUR_RUB)
}
