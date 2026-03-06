package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Калькулятор Go v1.0 ===")
	for {
		operacion := inputOperacion()
		number := inputNum()
		result := calculator(operacion, number)
		fmt.Println("Результат:", result)
		if checkRepeat() {
			continue
		}
		fmt.Println("До встречи ...")
		return
	}

}
func checkRepeat() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Хотите продолжить(y/n): ")
	scanner.Scan()

	if strings.ToLower(scanner.Text()) != "y" {
		return false
	}
	return true
}
func calculator(operacion string, number []float64) float64 {
	var resault float64
	n := len(number)
	switch operacion {
	case "SUM":
		for _, num := range number {
			resault += num
		}
		return resault
	case "AVG":
		for _, num := range number {
			resault += num
		}
		return resault / float64(n)
	case "MED":
		sort.Float64s(number)
		if n%2 != 0 {
			return number[n/2]
		}
		return (number[n/2-1] + number[n/2]) / 2
	}
	return 0
}

func inputOperacion() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Введите операцию (SUM, AVG, MED):")
		scanner.Scan()

		operacion := strings.ToUpper(strings.TrimSpace(scanner.Text()))

		if operacion == "SUM" || operacion == "AVG" || operacion == "MED" {
			return operacion
		}

		fmt.Printf("Ошибка: '%s' не поддерживается. Попробуйте еще раз.\n", operacion)

	}

}

func inputNum() []float64 {
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		flagError := false

		fmt.Println("Введите числа через пробел или запятую:")
		scanner.Scan()
		input = scanner.Text()

		if strings.TrimSpace(input) == "" {
			fmt.Println("Вы ничего не ввели")
			continue
		}

		parts := strings.FieldsFunc(input, func(r rune) bool {
			return r == ',' || r == '.' || r == ' '
		})
		numbers := make([]float64, 0, len(parts))

		for _, numStr := range parts {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Printf("Ошибка: %s не является числом !\n ", numStr)
				flagError = true
				break
			}
			numbers = append(numbers, num)
		}

		if flagError {
			fmt.Println("Попробуйте снова ...\n")
			continue
		}
		if len(numbers) == 0 {
			fmt.Println("Вы не ввели ни одного числа\n")
			continue
		}
		return numbers
	}
}
