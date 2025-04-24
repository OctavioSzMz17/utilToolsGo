package utilToolsGo

import (
	"fmt"
	"strconv"
	"math"
)

// ConvertToInt
func ConvertToInt(input string) (int, error) {
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("error al convertir el número: %v", err)
	}
	return number, nil
}

// This function adds two numbers and returns the result
func Add(a, b string) int {
	num1, _ := ConvertToInt(a)
	num2, _ := ConvertToInt(b)

	return num1 + num2
}

// This function less two numbers and returns the result
func Less(a, b string) int {
	num1, _ := ConvertToInt(a)
	num2, _ := ConvertToInt(b)

	return num1 - num2
}

// This function multiply two numbers and returns the result
func Multiply(a, b string) int {
	num1, _ := ConvertToInt(a)
	num2, _ := ConvertToInt(b)

	return num1 * num2
}

// This function divide two numbers and returns the result
func Divide(a, b string) int {
	num1, _ := ConvertToInt(a)
	num2, _ := ConvertToInt(b)

	return num1 / num2
}

func Sqrt(a string) int {
	num1, _ := ConvertToInt(a)

	return int(math.Sqrt(float64(num1)))
}

func Pow(a, b string) int {
	num1, _ := ConvertToInt(a)
	num2, _ := ConvertToInt(b)	
	
	return int(math.Pow(float64(num1), float64(num2)))
}