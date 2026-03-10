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
	for {
		amount, from, to := readInput()

		result := converterMap(amount, from, to)
		fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)

		fmt.Println("\n--- Курсы валют ---")
		fmt.Printf("Курс USD -> EUR: %.2f\n", usdToEur)
		fmt.Printf("Курс USD -> RUB: %.2f\n", usdToRub)
		fmt.Println("-------------------")
		fmt.Printf("Рассчитанный курс EUR -> RUB: %.2f\n", eurToRub)

		isCheckRepeat := checkRepeat()
		if isCheckRepeat {
			continue
		}
		fmt.Println("До встречи ...")
		return

	}
}

var converter = map[string]float64{
	"usd": 1.0,
	"eur": 0.84,
	"rub": 76.75,
}

func converterMap(amount float64, from, to string) float64 {
	converterFrom := converter[from]
	converterTo := converter[to]

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

// func convert(amount float64, from, to string) float64 {

// 	switch from {
// 	case "usd":
// 		if to == "eur" {
// 			return amount * usdToEur
// 		}
// 		if to == "rub" {
// 			return amount * usdToRub
// 		}

// 	case "eur":
// 		if to == "usd" {
// 			return amount / usdToEur
// 		}
// 		if to == "rub" {
// 			return amount * (1 / usdToEur) * usdToRub
// 		}

// 	case "rub":
// 		if to == "usd" {
// 			return amount / usdToRub
// 		}
// 		if to == "eur" {
// 			return amount / ((1 / usdToEur) * usdToRub)
// 		}
// 	}
// 	return 0
// }

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
