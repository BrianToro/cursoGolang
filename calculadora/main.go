package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(num1 int, num2 int, number string) (int, bool) {
	var result int
	var error bool = false
	switch number {
	case "+":
		result = num1 + num2
		break
	case "-":
		result = num1 - num2
		break
	case "*", "x":
		result = num1 * num2
		break
	case "/":
		result = num1 / num2
		break
	default:
		error = true
	}
	return result, error
}

func main() {
	rawInput := readInput()
	values := processInput(rawInput)
	number1, err1 := strconv.Atoi(values[0])
	number2, err2 := strconv.Atoi(values[2])
	processResult(number1, number2, err1, err2, values[1])
}

func processResult(number1 int, number2 int, err1 error, err2 error, operator string) {
	c := calc{}
	if (err1 != nil || err2 != nil) || number2 == 0 {
		errorController()
	} else {
		result, error := c.operate(number1, number2, operator)
		if error {
			errorController()
		} else {
			fmt.Println("Result of calc: ", result)
		}
	}
}

func errorController() {
	fmt.Println("----------- Invalid operation -----------")
}

func processInput(rawInput string) []string {
	values := strings.Split(rawInput, " ")
	return values
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
