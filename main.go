package main

import (
	"fmt"
	"strings"
)

const (
	usdToEur = 0.84
	usdToRub = 76.75
	eurToRub = (1 / usdToEur) * usdToRub
)

func main() {
	converter := map[string]float64{
		"usd": 1.0,
		"eur": 0.84,
		"rub": 76.75,
	}

	for {
		amount, from, to := readInput()

		result := converterMap(amount, from, to, &converter)
		fmt.Println("ВАШ РАССЧЕТ: ")
		fmt.Printf("\n%.2f %s = %.2f %s\n", amount, from, result, to)

		showrates(&converter)
		isCheckRepeat := checkRepeat()
		if isCheckRepeat {
			continue
		}
		fmt.Println("До встречи ...")
		return

	}
}
func showrates(converter *map[string]float64) {
	fmt.Println("\n--- Курсы валют ---")
	fmt.Printf("Курс USD -> EUR: %.2f\n", usdToEur)
	fmt.Printf("Курс USD -> RUB: %.2f\n", usdToRub)
	fmt.Println("-------------------")
	fmt.Printf("Рассчитанный курс EUR -> RUB: %.2f\n", eurToRub)
	fmt.Println("-------------------")
	fmt.Println("\n--- Текущие курсы ---")
	for curr, rate := range *converter {
		fmt.Printf("%s: %.2f\n", strings.ToUpper(curr), rate)
	}
	fmt.Println("-------------------")
}

func converterMap(amount float64, from, to string, converter *map[string]float64) float64 {

	converterFrom := (*converter)[from]
	converterTo := (*converter)[to]

	if converterFrom == 0 {
		return 0
	}

	return amount * (converterTo / converterFrom)

}
func checkRepeat() bool {
	var vibor string
	fmt.Println("Хотите продолжить ? нажмите  (y/n): ")
	fmt.Scan(&vibor)
	lowerVibor := strings.ToLower(vibor)
	if lowerVibor == "y" {
		return true
	}
	return false
}

func readInput() (float64, string, string) {

	from := inputFrom()
	amount := inputAmount()
	to := inputTo(from)
	return amount, from, to
}

func inputTo(from string) string {
	var to string
	for {
		fmt.Printf("Введите целевую валюту (кроме %s): ", strings.ToUpper(from))
		fmt.Scan(&to)
		lowerTo := strings.ToLower(to)
		if lowerTo == from {
			fmt.Println("Целевая валюта не должна совпадать с исходной!")
			continue
		}
		if lowerTo == "usd" || lowerTo == "eur" || lowerTo == "rub" {
			return lowerTo
		}
		fmt.Println("Введите праивльное название валюты !")
		continue
	}
}

func inputAmount() float64 {
	var amount float64
	for {
		fmt.Println("Введите сумму [НЕ ОТРИЦАТЕЛЬНО ЧИСЛО]: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Число не может быть меньше или равно 0")
			continue
		}
		return float64(amount)
	}

}

func inputFrom() string {
	var from string
	for {
		fmt.Print("Введите исходную валюту (например USD/EUR/RUB): ")
		fmt.Scan(&from)
		lowerFrom := strings.ToLower(from)
		if lowerFrom == "usd" || lowerFrom == "eur" || lowerFrom == "rub" {
			return lowerFrom
		}
		fmt.Println("Введите праивльное название валюты !")
		continue
	}

}
