package main

import (
	"os"
)

func IsValidOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/" || op == "%"
}

func Atoi(str string) (int, bool) {
	if len(str) == 0 {
		return 0, false
	}
	result := 0
	sign := 1
	start := 0
	if str[0] == '-' {
		sign = -1
		start = 1
	} else if str[0] == '+' {
		start = 1
	}
	if start >= len(str) {
		return 0, false
	}
	for i := start; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return 0, false
		}
		digit := int(str[i] - '0')
		if result > (9223372036854775807-digit)/10 {
			return 0, false
		}
		result = result*10 + digit
	}
	return result * sign, true
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	result := ""
	for n > 0 {
		result = string(rune(n%10+'0')) + result
		n /= 10
	}
	if negative {
		result = "-" + result
	}
	return result
}

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}
	num1Str := args[1]
	operator := args[2]
	num2Str := args[3]
	if !IsValidOperator(operator) {
		return
	}
	num1, ok1 := Atoi(num1Str)
	if !ok1 {
		return
	}
	num2, ok2 := Atoi(num2Str)
	if !ok2 {
		return
	}
	if num2 == 0 {
		if operator == "/" {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		if operator == "%" {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
	}
	var result int
	var overflow bool
	switch operator {
	case "+":
		result = num1 + num2
		if (num1 > 0 && num2 > 0 && result < 0) || (num1 < 0 && num2 < 0 && result > 0) {
			overflow = true
		}
	case "-":
		result = num1 - num2
		if (num1 > 0 && num2 < 0 && result < 0) || (num1 < 0 && num2 > 0 && result > 0) {
			overflow = true
		}
	case "*":
		result = num1 * num2
		if num1 != 0 && result/num1 != num2 {
			overflow = true
		}
	case "/":
		result = num1 / num2
	case "%":
		result = num1 % num2
	}
	if overflow {
		return
	}
	os.Stdout.Write([]byte(Itoa(result) + "\n"))
}
