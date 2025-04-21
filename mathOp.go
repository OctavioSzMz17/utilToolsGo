package utilToolsGo

import (
	"fmt"
	"strconv"
)

// ConvertToInt
func convertToInt(input string) (int, error) {
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Error al convertir el número: %v", err)
	}
	return number, nil
}

// This function adds two numbers and returns the result
func Add(a, b string) (int, error) {
	num1, _ := convertToInt(a)
	num2, _ := convertToInt(b)

	return num1 + num2, nil
}

// This function less two numbers and returns the result
func Less(a, b string) (int, error) {
	num1, _ := convertToInt(a)
	num2, _ := convertToInt(b)

	return num1 - num2, nil
}

// This function multiply two numbers and returns the result
func Multiply(a, b string) (int, error) {
	num1, _ := convertToInt(a)
	num2, _ := convertToInt(b)

	return num1 * num2, nil
}

// This function divide two numbers and returns the result
func Divide(a, b string) (int, error) {
	num1, _ := convertToInt(a)
	num2, _ := convertToInt(b)

	if num2 == 0 {
		return 0, fmt.Errorf("Error: Division by zero")
	}
	return num1 / num2, nil
}
