package main

import "fmt"

func main() {
	const (
		usdToEur = 0.84
		usdToRub = 76.75
	)

	const eurToRub = (1 / usdToEur) * usdToRub

	fmt.Println("--- Курсы валют ---")
	fmt.Printf("Курс USD -> EUR: %.2f\n", usdToEur)
	fmt.Printf("Курс USD -> RUB: %.2f\n", usdToRub)
	fmt.Println("-------------------")
	fmt.Printf("Рассчитанный курс EUR -> RUB: %.2f\n", eurToRub)
}
