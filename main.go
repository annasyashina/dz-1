package main

import (
	"fmt"
)

type convertMap = map[string]float64

const USD_EUR = 0.8599
const USD_RUB = 78.84

func main() {

	var sourceCurrency string
	var wantConvet float64
	var convertibleCurrency string
	sourceCurrency, wantConvet, convertibleCurrency = getUserInput()
	//fmt.Println(sourceCurrency, wantConvet, convertibleCurrency)
	m_evro := make(convertMap, 2)
	m_evro["USD"] = 1 / USD_EUR
	m_evro["RUB"] = USD_RUB / USD_EUR

	m_rub := make(convertMap, 2)
	m_rub["USD"] = 1 / USD_RUB
	m_rub["EUR"] = USD_EUR / USD_RUB

	m_usd := make(convertMap, 2)
	m_usd["RUB"] = USD_RUB
	m_usd["EUR"] = USD_EUR

	m := map[string]convertMap{"USD": m_usd, "EUR": m_evro, "RUB": m_rub}
	result := calculation(m, wantConvet, sourceCurrency, convertibleCurrency)
	//fmt.Println(m[sourceCurrency][convertibleCurrency])
	//fmt.Println(reflect.TypeOf((m[sourceCurrency][convertibleCurrency])))
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

func calculation(m map[string]convertMap, number float64, current string, convert string) float64 {
	var result float64

	//fmt.Println(*m[current])
	result = (number) * (m[current][convert])
	/*	switch current {
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

		}*/

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
