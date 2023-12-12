package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transformationIntoArabian(str string) int {
	numbers := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return numbers[str]
}

func transformationIntoRoman(integer int) string {
	result := ""
	for integer > 0 {
		if (integer - 500) >= 0 {
			result += "D"
			integer -= 500
			continue
		}
		if (integer - 400) >= 0 {
			result += "CD"
			integer -= 400
			continue
		}
		if (integer - 100) >= 0 {
			result += "C"
			integer -= 100
			continue
		}
		if (integer - 90) >= 0 {
			result += "XC"
			integer -= 90
			continue
		}
		if (integer - 50) >= 0 {
			result += "L"
			integer -= 50
			continue
		}
		if (integer - 40) >= 0 {
			result += "XL"
			integer -= 40
			continue
		}
		if (integer - 10) >= 0 {
			result += "X"
			integer -= 10
			continue
		}
		if (integer - 9) >= 0 {
			result += "IX"
			integer -= 9
			continue
		}
		if (integer - 5) >= 0 {
			result += "V"
			integer -= 5
			continue
		}
		if (integer - 4) >= 0 {
			result += "IV"
			integer -= 4
			continue
		}
		if (integer - 1) >= 0 {
			result += "I"
			integer -= 1
			continue
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		operands := strings.Split(text, " ")
		if len(operands) < 3 {
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			break
		} else if len(operands) > 3 {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			break
		}

		symbol := operands[1]

		arg1, _ := strconv.Atoi(operands[0])
		arg2, _ := strconv.Atoi(operands[2])
		RimArg1 := transformationIntoArabian(operands[0])
		RimArg2 := transformationIntoArabian(operands[2])

		if (arg1 < 1 && RimArg1 == 0 && RimArg2 == 0) || arg1 > 10 || (arg2 < 1 && RimArg1 == 0 && RimArg2 == 0) || arg2 > 10 {
			fmt.Println("Вывод ошибки, так как калькулятор принимает на вход числа от 1 до 10 включительно")
			break
		}

		if arg1 == 0 && arg2 == 0 && RimArg1 != 0 && RimArg2 != 0 {
			switch symbol {
			case "+":
				fmt.Println(transformationIntoRoman(RimArg1 + RimArg2))
			case "-":
				if (RimArg1 - RimArg2) > 0 {
					fmt.Println(transformationIntoRoman(RimArg1 - RimArg2))
				} else {
					fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
				}
			case "/":
				fmt.Println(transformationIntoRoman(RimArg1 / RimArg2))
			case "*":
				fmt.Println(transformationIntoRoman(RimArg1 * RimArg2))
			default:
				fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			}
		} else if arg1 != 0 && arg2 != 0 && RimArg1 == 0 && RimArg2 == 0 {
			if arg1 < 1 || arg1 > 10 || arg2 < 1 || arg2 > 10 {
				fmt.Println("Вывод ошибки, так как калькулятор принимает на вход числа от 1 до 10 включительно")
				break
			}

			switch symbol {
			case "+":
				fmt.Println(arg1 + arg2)
			case "-":
				fmt.Println(arg1 - arg2)
			case "/":
				fmt.Println(arg1 / arg2)
			case "*":
				fmt.Println(arg1 * arg2)
			default:
				fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			}
		} else {
			fmt.Println("Вывод ошибки, из-за некорректного ввода")
			break
		}

	}
}
