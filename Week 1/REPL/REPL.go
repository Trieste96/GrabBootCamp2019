package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//eval
func eval(input string) (result string, err error) {
	s := strings.Split(strings.Trim(input, "\n"), " ")
	if len(s) < 2 {
		return "", errors.New("Invalid format!")
	}
	num1, err1 := strconv.ParseFloat(s[0], 64)
	num2, err2 := strconv.ParseFloat(s[2], 64)
	op := s[1]

	if err1 == nil && err2 == nil {
		result := 0.0
		switch op {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				return "", errors.New("Divided by zero!")
			}
			result = num1 / num2
		default:
			return "", errors.New("Unrecognized operator!")
		}
		return fmt.Sprint("> ", num1, " ", op, " ", num2, " = ", result), nil
	} else {
		return "", errors.New("ERROR!")
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		result, err := eval(input)
		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	}

}
