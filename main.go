package main

import "fmt"

func main() {

	/*var sourceCurrency string
	var wantConvet string
	var convertibleCurrency string
	sourceCurrency, wantConvet, convertibleCurrency = getUserInput()
	*/
	const USD_EUR = 0.8549
	const USD_RUB = 81.9
	//EUR_USD := 1 / USD_EUR
	//RUB_USD := 1 / USD_RUB
	//EUR_RUB := EUR_USD / RUB_USD
	EUR_RUB := USD_RUB / USD_EUR
	fmt.Println(EUR_RUB)
}

func getUserInput() (string, string, string) {
	var sourceCurrency string
	var wantConvet string
	var convertibleCurrency string
	fmt.Println("Введите исходную валюту:")
	fmt.Scan(&sourceCurrency)
	fmt.Println("Введите, что хотим конвертировать:")
	fmt.Println(&wantConvet)
	fmt.Println("Введите конвертируемую валюту:")
	fmt.Scan(&convertibleCurrency)
	return sourceCurrency, wantConvet, convertibleCurrency
}

func calculation(float64, string, string) {

}
