package main

import (
	"fmt"
)

const USD_EUR = 0.8599
const USD_RUB = 78.84

func main() {

	var sourceCurrency string
	var wantConvet float64
	var convertibleCurrency string
	sourceCurrency, wantConvet, convertibleCurrency = getUserInput()
	fmt.Println(sourceCurrency, wantConvet, convertibleCurrency)

	result := calculation(sourceCurrency, wantConvet, convertibleCurrency)
	fmt.Println(result)
	//EUR_USD := 1 / USD_EUR
	//RUB_USD := 1 / USD_RUB
	//EUR_RUB := EUR_USD / RUB_USD
	//EUR_RUB := USD_RUB / USD_EUR
	//fmt.Println(EUR_RUB)
}

func getUserInput() (string, float64, string) {
	var sourceCurrency string
	var number float64
	var convertibleCurrency string

	currency := []string{}
	currency = append(currency, "USD")
	currency = append(currency, "EUR")
	currency = append(currency, "RUB")

	isCurrency := false

	for isCurrency == false {
		fmt.Println("Введите исходную валюту", currency, ":")
		fmt.Scan(&sourceCurrency)
		if elementInArray(sourceCurrency, currency) {
			isCurrency = true
		}
	}

	isFloat := false
	for !isFloat {
		fmt.Println("Введите, что хотим конвертировать:")
		fmt.Scan(&number)
		if number > 0.0 || number > 0 {
			isFloat = true
		}
	}

	convert := []string{}
	for i := 0; i < len(currency); i++ {
		if currency[i] != sourceCurrency {
			convert = append(convert, currency[i])
		}
	}

	isCurrency = false

	for isCurrency == false {
		fmt.Println("Введите конвертируемую валюту", convert, ":")
		fmt.Scan(&convertibleCurrency)
		if elementInArray(convertibleCurrency, convert) {
			isCurrency = true
		}
	}

	//fmt.Println("Введите конвертируемую валюту:")
	//fmt.Scan(&convertibleCurrency)
	return sourceCurrency, number, convertibleCurrency
}

func calculation(current string, number float64, convert string) float64 {
	var result float64

	switch current {
	case "USD":
		switch convert {
		case "RUB":
			result = number * USD_RUB
		case "EUR":
			result = number * USD_EUR
		}

	case "RUB":
		switch convert {
		case "USD":
			result = number / USD_RUB
		case "EUR":
			result = number * USD_EUR / USD_RUB
		}

	case "EUR":
		switch convert {
		case "RUB":
			result = number * USD_RUB / USD_EUR
		case "USD":
			result = number / USD_EUR
		}

	}

	return result
}

func elementInArray(element string, elements []string) bool {
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return true
		}
	}
	return false
}
