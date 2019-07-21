package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validate(input string) (num1, num2 float64, op string, err error) {
	s := strings.Fields(input)
	if len(s) != 3 {
		return 0, 0, "", errors.New("Invalid input")
	}

	num1, err = strconv.ParseFloat(s[0], 64)
	if err != nil {
		return 0, 0, "", errors.New("invalid format on first operand ")
	}

	num2, err = strconv.ParseFloat(s[1], 64)
	if err != nil {
		return 0, 0, "", errors.New("invalid format on second operand ")
	}

	op = s[1]
	if strings.ContainsAny(op, "+-/*") == false {
		return 0, 0, "", errors.New("invalid operator")
	}
	return num1, num2, op, nil
}

//eval
func eval(input string) (result string, err error) {
	num1, num2, operator, err := validate(input)

	if err != nil {
		return "", err
	}
	var res float64
	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("divided by zero")
		}
		res = num1 / num2
	default:
		return "", errors.New("unrecognized operator")
	}
	return fmt.Sprint("> ", num1, " ", operator, " ", num2, " = ", res), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}

		result, err := eval(input)
		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	}

}
