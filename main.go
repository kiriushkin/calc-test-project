package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var vocab = map[string]int {
	"M": 1000,
	"CM": 900,
	"D": 500,
	"CD": 400,
	"C": 100,
	"XC": 90,
	"L": 50,
	"XL": 40,
  "X": 10,
	"IX": 9,
	"V": 5,
	"IV": 4,
	"I": 1,
}

var conversions = []struct {
		value int
		digit string
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	slice := strings.Split(text, " ")

	if len(slice) != 3 {
		panic("Неверый формат математической операции")
	}

	var a, b int

	first := slice[0]
	op := slice[1]
	second := slice[2]

	aType := IsRomanOrArabic(first)
	bType := IsRomanOrArabic(second)

	if aType == "Neither" || bType == "Neither" {
		panic("Число(-а) не являются подходящими")
	}

	if aType != bType {
		panic("Используются разные системы счисления")
	}

	if aType == "Roman" {
		a = RomanToInt(first)
		b = RomanToInt(second)
	} else {
		a, _ = strconv.Atoi(first)
		b, _ = strconv.Atoi(second)
	}

	if a > 10 || a < 1 || b > 10 || b < 1 {
		panic("Число(-а) выходит(-ят) за пределы допустимого диапазона")
	}

	result, err := Calculate(a, b, op)

	if err {
		panic("Некорректный оператор")
	}

	if aType == "Roman" {
		if (result < 1) {
			panic("Отрицательный результат или нулевой для римских чисел")
		}
		strResult := IntToRoman(result)
		fmt.Println(strResult)
		return
	}

	fmt.Println(result)
}

func IsRomanOrArabic(num string) string {
	if RomanToInt(num) != 0 {
		return "Roman"
	}
	if _, err := strconv.Atoi(num); err == nil {
		return "Arabic"
	}
	return "Neither"
}

func RomanToInt(s string) int {
	length := len(s)
	lastEl := s[length-1 : length]
	var result int
	result = vocab[lastEl]
	for i := length - 1; i > 0; i-- {
		if vocab[s[i:i+1]] <= vocab[s[i-1:i]] {
			result += vocab[s[i-1:i]]
		} else {
			result -= vocab[s[i-1:i]]
		}
	}
	return result
}

func IntToRoman(n int) string {
	var resultArr []string

	for _, conversion := range conversions {
		for n >= conversion.value {
			resultArr = append(resultArr, conversion.digit)
			n -= conversion.value
		}
	}

	return strings.Join(resultArr, "")
}

func Calculate(a, b int, op string) (int, bool) {
	switch op {
	case "+": return a + b, false
	case "-": return a - b, false
	case "*": return a * b, false
	case "/": return a / b, false
	default: return 0, true
	}
}