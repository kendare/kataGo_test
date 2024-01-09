package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isArabicNumber(num string) bool {
	n, err := strconv.Atoi(num)
	return err == nil && n >= 1 && n <= 10
}

var romanNumerals = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
}

var arabicNumerals = []struct {
	arabic int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
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

func romanToArabic(roman string) (int, error) {
	result := 0
	roman = strings.ToUpper(roman)

	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && romanNumerals[roman[i:i+2]] > 0 {
			result += romanNumerals[roman[i:i+2]]
			i++
		} else {
			result += romanNumerals[roman[i:i+1]]
		}
	}

	if result == 0 {
		return 0, fmt.Errorf("неверное римское число: %s", roman)
	}

	return result, nil
}

func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 || arabic > 100 {
		err := fmt.Errorf("Результат меньше нуля не поддерживается в римской системе счисления")
		return "", err
	}

	result := ""
	for _, numeral := range arabicNumerals {
		for arabic >= numeral.arabic {
			result += numeral.roman
			arabic -= numeral.arabic
		}
	}

	return result, nil
}

func isRomanNumber(num string) bool {
	romanNumerals := map[string]bool{
		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
	}
	_, found := romanNumerals[num]
	return found
}

//func romanToArabic(roman string) (int, error) {
// реализация преобразования римского числа в арабское
//}

func calculate(arabicA, arabicB int, operator string) (int, error) {
	var res int
	switch operator {
	case "+":
		res = arabicA + arabicB
	case "-":
		res = arabicA - arabicB
	case "*":
		res = arabicA * arabicB
	case "/":
		res = arabicA / arabicB
	}
	return res, nil
}

func isValidNumber(num string) bool {
	return isArabicNumber(num) || isRomanNumber(num)
}

// ... (остальной код без изменений)

func main() {
	fmt.Print("Введите выражение: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Fields(input)

	if len(parts) != 3 {
		fmt.Println("Ошибка: Неправильный формат ввода. Используйте два операнда и один оператор (+, -, *, /).")
		os.Exit(1)
	}

	operandA, operandB, operator := parts[0], parts[2], parts[1]

	if (isArabicNumber(operandA) && isRomanNumber(operandB)) || (isRomanNumber(operandA) && isArabicNumber(operandB)) {
		fmt.Println("Ошибка: Используются одновременно разные системы счисления.")
		os.Exit(1)
	}
	if !isValidNumber(operandA) || !isValidNumber(operandB) {
		fmt.Println("Ошибка: Числа должны быть в диапазоне от 1 до 10 включительно")
		os.Exit(1)
	}
	var result int
	var err error
	var res string
	if isArabicNumber(operandA) {
		arabicA, _ := strconv.Atoi(operandA)
		arabicB, _ := strconv.Atoi(operandB)
		result, err = calculate(arabicA, arabicB, operator)
		res = strconv.Itoa(result)
	} else if isRomanNumber(operandA) {
		romanA := operandA
		romanB := operandB
		arabicA, _ := romanToArabic(romanA)
		arabicB, _ := romanToArabic(romanB)
		result, err = calculate(arabicA, arabicB, operator)
		res, err = arabicToRoman(result)
	}

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}
