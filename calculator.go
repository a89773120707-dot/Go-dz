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
		nums := inputNum()
		result := calculator(operacion, nums)
		fmt.Println("Результат:", result)
		if checkRepeat() {
			continue
		}
		fmt.Println("До встречи ...")
		return
	}

}

func calculator(operacion string, nums []float64) float64 {
	type Calculation func(nums []float64) float64
	var resault float64
	n := len(nums)
	menuOperation := map[string]Calculation{
		"SUM": func(nums []float64) float64 {
			for _, num := range nums {
				resault += num
			}
			return resault
		},
		"AVG": func(nums []float64) float64 {
			for _, num := range nums {
				resault += num
			}
			return resault / float64(n)
		},
		"MED": func(nums []float64) float64 {
			n := len(nums)
			if n == 0 {
				return 0
			}
			sort.Float64s(nums)
			if n%2 != 0 {
				return nums[n/2]
			}
			return (nums[n/2-1] + nums[n/2]) / 2
		},
	}
	return menuOperation[operacion](nums)
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
			return r == ',' || r == ' '
		})
		numss := make([]float64, 0, len(parts))

		for _, numStr := range parts {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Printf("Ошибка: %s не является числом !\n ", numStr)
				flagError = true
				break
			}
			numss = append(numss, num)
		}

		if flagError {
			fmt.Println("Попробуйте снова ...\n")
			continue
		}
		if len(numss) == 0 {
			fmt.Println("Вы не ввели ни одного числа\n")
			continue
		}
		return numss
	}
}
