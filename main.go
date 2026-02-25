package main

import "fmt"

func main() {
	const (
		usdToEur = 0.84
		usdToRub = 76.75
	)

	const eurToRub = (1 / usdToEur) * usdToRub

	amount, from, to := readInput()
	fmt.Println(amount, from, " -> ", to)

	fmt.Println("--- Курсы валют ---")
	fmt.Printf("Курс USD -> EUR: %.2f\n", usdToEur)
	fmt.Printf("Курс USD -> RUB: %.2f\n", usdToRub)
	fmt.Println("-------------------")
	fmt.Printf("Рассчитанный курс EUR -> RUB: %.2f\n", eurToRub)
}

func readInput() (float64, string, string) {
	var amount float64
	var from string
	var to string

	fmt.Print("Введите сумму : ")
	fmt.Scan(&amount)

	fmt.Print("Введите исходную валюту (например USD): ")
	fmt.Scan(&from)

	fmt.Print("Введите целевую валюту (например EUR): ")
	fmt.Scan(&to)

	return amount, from, to
}
