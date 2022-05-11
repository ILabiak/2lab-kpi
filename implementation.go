package lab2 //package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// TODO: document this function.
// Calculate postfix expression
func CalculatePostfix(input string) (string, error) {
	var stack []float64
	var chars = strings.Fields(input)
	var value float64
	for _, char := range chars {
		switch char {
		case "+", "-", "*", "/", "^":
			var firstVal, secondVal float64
			if len(stack) < 2 {
				return "0", fmt.Errorf("There are no 2 values, can't calculate expression - %s", char)
			}
			stack, firstVal, secondVal = GetValuesFromStack(stack)
			switch char {
			case "+":
				value = firstVal + secondVal
			case "-":
				value = firstVal - secondVal
			case "*":
				value = firstVal * secondVal
			case "/":
				value = firstVal / secondVal
			case "^":
				value = math.Pow(firstVal, secondVal)
			}
		default:
			var err error
			if value, err = strconv.ParseFloat(char, 64); err != nil {
				return "0", fmt.Errorf("Wrong value - %s", char)
			}
		}
		stack = append(stack, value)
	}
	if len(stack) != 1 {
		return "0", fmt.Errorf("Wrong expression - %s", stack)
	}
	str := fmt.Sprintf("%f", stack[0])
	return str, nil
}


func GetValuesFromStack(stack []float64) ([]float64, float64, float64) {
	var values []float64
	stack, values = stack[:len(stack)-2], stack[len(stack)-2:]
	//fmt.Println(values[0], values[1])
	return stack, values[0], values[1]
}
