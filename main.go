package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a, b *int

var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

const (
	NEGVAL = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	DIFVAL = "Вывод ошибки, так как используются одновременно разные системы счисления."
	NOMATH = "Вывод ошибки, так как строка не является математической операцией."
	FORMAT = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	NULL   = "Вывод ошибки, так как в римской системе счислений нет значения равное 0"
	INCORR = "Вывод ошибки, так как использованы числа не из диапазона [1:10]"
)

var romanNum1, romanNum2 int

func getRomanNum(romanResult int) string {
	var romanNum string
	if romanResult == 0 {
		panic(NULL)
	} else if romanResult < 0 {
		panic(NEGVAL)
	}
	for romanResult > 0 {
		for _, val := range convIntToRoman {
			for i := val; i <= romanResult; {
				for index, value := range roman {
					if value == val {
						romanNum += index
						romanResult -= val
					}
				}
			}
		}
	}
	return romanNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(strings.TrimSpace(input))
	c := strings.Split(input, " ")
	if len(c) != 3 {
		panic(NOMATH)
	}
	num1, err := strconv.Atoi(c[0])
	var stringFound int
	if err != nil {
		romanNum1 = roman[c[0]]
		stringFound++
	}
	num2, err := strconv.Atoi(c[2])
	if err != nil {
		romanNum2 = roman[c[2]]
		stringFound++
	}
	operator := c[1]
	result, truth := operators[operator]
	if !truth {
		panic(FORMAT)
	}
	switch stringFound {
	case 1:
		panic(DIFVAL)
	case 0:
		if num1 > 0 && num1 < 11 && num2 > 0 && num2 < 11 {
			a, b = &num1, &num2
			fmt.Println(result())
		} else {
			panic(INCORR)
		}
	case 2:
		if romanNum1 > 0 && romanNum1 < 11 && romanNum2 > 0 && romanNum2 < 11 {
			a, b = &romanNum1, &romanNum2
			fmt.Println(getRomanNum(result()))
		} else {
			panic(INCORR)
		}
	}
}
