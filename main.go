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
	fmt.Println("Калькулятор 1.0")
	for {
		op, data := collectUserData()
		result := calculate(op, data)
		fmt.Println("Ваша операция", op, "Результат:", result)
		fmt.Println("Вы хотите сделать еще расчёт?")
		answer := ""
		_, _ = fmt.Scan(&answer)
		if answer == "Да" || answer == "да" {
			continue
		} else {
			break
		}
	}
}

func collectUserData() (operation string, digits []int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите операцию (AVG, SUM, MED): ")
	for {
		scanner.Scan()
		text := scanner.Text()
		if text != "AVG" && text != "SUM" && text != "MED" {
			fmt.Println("Неверные данные, введите операцию (AVG, SUM, MED): ")
			continue
		} else {
			operation = strings.Trim(scanner.Text(), " ")
			break
		}
	}
	fmt.Println("Введите числа (через запятую): ")
	scanner.Scan()
	input := scanner.Text()
	input = strings.ReplaceAll(input, ",", " ")
	digitsStr := strings.Fields(input)
	digits = make([]int, len(digitsStr))
	for i := range digitsStr {
		num, err := strconv.Atoi(digitsStr[i])
		if err != nil {
			fmt.Println(err)
		}
		digits[i] = num
	}
	return operation, digits
}

func calculate(operation string, digits []int) (result float64) {
	switch operation {
	case "AVG":
		sum := 0
		if len(digits) == 0 {
			return 0
		}
		for _, d := range digits {
			sum += d
		}
		result = float64(sum) / float64(len(digits))
	case "SUM":
		sum := 0
		for _, d := range digits {
			sum += d
		}
		result = float64(sum)
	case "MED":
		n := len(digits)
		if n == 0 {
			return 0
		}
		sorted := make([]int, n)
		copy(sorted, digits)
		sort.Ints(sorted)

		mid := n / 2
		if n%2 == 0 {
			return float64(sorted[mid-1]+sorted[mid]) / 2.0
		} else {
			return float64(sorted[mid])
		}
	}

	return result
}
