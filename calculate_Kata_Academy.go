package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func romanToArabic(r string) int {
	arabicNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
	}
	result := 0
	preValue := 0

	for _, char := range r {
		value := arabicNumerals[char]
		if value > preValue {
			result += value - 2*preValue // Если текущее число больше предыдущего, вычитаем дважды предыдущее
		} else {
			result += value
		}
		preValue = value
	}

	return result
}

func ArabicToRoman(arabic int) string {
	type romanNumerals struct {
		Value  int
		Symbol string
	}
	var allRomanNumeruls = []romanNumerals{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	result := ""
	for _, num := range allRomanNumeruls {
		for arabic >= num.Value {
			result += num.Symbol
			arabic -= num.Value
		}

	}
	return result
}

func calculate(a, c int, b string) int {
	switch b {
	case "+":
		return a + c
	case "-":
		return a - c
	case "*":
		return a * c
	case "/":
		return a / c
	}
	return 0
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	name := scanner.Text()
	number := strings.Fields(name)

	if len(number) != 3 {
		panic("cтрока не является математической операцией")
	} else {

		sliceAC := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		sliceB := []string{"+", "-", "/", "*"}
		sliceRome := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

		if contains(sliceAC, number[0]) && contains(sliceAC, number[2]) {
			if contains(sliceB, number[1]) {
				inta, _ := strconv.Atoi(number[0])
				intc, _ := strconv.Atoi(number[2])

				resultArabic := calculate(inta, intc, number[1])

				fmt.Println(resultArabic)
			} else {
				panic("неизвестная арифметическая операция")
			}

		} else {
			if contains(sliceRome, number[0]) && contains(sliceRome, number[2]) {
				if contains(sliceB, number[1]) {
					arabic1 := romanToArabic(number[0])
					arabic2 := romanToArabic(number[2])

					result := calculate(arabic1, arabic2, number[1])
					if result <= 0 {
						panic("в римской системе нет отрицательных чисел")
					} else {
						resultRoman := ArabicToRoman(result)
						fmt.Println(resultRoman)
					}
				} else {
					panic("неизвестная арифметическая операция")
				}

			} else {
				panic("используются одновременно разные системы счисления")
			}
		}
	}
}
