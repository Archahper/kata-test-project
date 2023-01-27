package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanDigits = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	var (
		input, operator, strNum1, strNum2 string
		num1, num2, result                int
		isRoman                           bool
	)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		fmt.Println("Формат ввода: Операнд1 Оператор Операнд2")
		return
	}

	words := strings.Split(input, " ")
	if len(words) != 3 {
		fmt.Println("Формат ввода: Операнд1 Оператор Операнд2")
		return
	}

	strNum1 = words[0]
	operator = words[1]
	strNum2 = words[2]

	num1, err1 := romanToInt(strNum1)
	num2, err2 := romanToInt(strNum2)

	if err1 == nil && err2 == nil {
		isRoman = true
	} else if (err1 != nil) != (err2 != nil) {
		fmt.Println("Цифры должны быть либо римские либо арабские. Римские цифры могут быть только положительные.")
		return
	}

	if !isRoman {
		num1, err = strconv.Atoi(strNum1)
		num2, err = strconv.Atoi(strNum2)
		if err != nil {
			fmt.Println("Введеные числа не являются арабскими или римскими.")
			return
		}
	}

	if num1 < 1 || num2 < 1 || num1 > 10 || num2 > 10 {
		fmt.Println("Калькулятор работает с числами от 0 до 10 включительно!")
		return
	}

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Поддерживаемые операции: + - * /")
		return
	}

	if isRoman {
		if result < 1 {
			fmt.Println("Результатом операций с римскими числами может быть только положительное число!")
			return
		}
		resultRoman, _ := intToRoman(result)
		fmt.Println(resultRoman)
	} else {
		fmt.Println(result)
	}
}

func romanToInt(s string) (int, error) {
	var result, prevRoman int

	for i := 0; i < len(s); i++ {
		roman, ok := romanDigits[string(s[i])]
		if !ok {
			return result, errors.New("invalid syntax")
		}
		if i > 0 {
			prevRoman = romanDigits[string(s[i-1])]
		}
		if i > 0 && roman > prevRoman {
			result += roman - 2*prevRoman
		} else {
			result += roman
		}
	}
	return result, nil
}

func intToRoman(num int) (string, error) {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result string
	for i := 0; num > 0; i++ {
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}
	return result, nil
}
