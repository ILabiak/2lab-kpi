package lab2 //package lab2

import (
	"fmt"
	//"go/printer"
	"math"
	"strconv"
	"strings"
)

// This function accepts postfix expression as string and calculates it.
// Returns result as string and error.
// If there is no result - returns "Nil" and error.
// If there is no error - returns result and nil as error.
func CalculatePostfix(input string) (string, error) {
	var stack []float64
	var chars = strings.Fields(input)
	var value float64
	for _, char := range chars {
		switch char {
		case "+", "-", "*", "/", "^":
			var firstVal, secondVal float64
			if len(stack) < 2 {
				return "Nil", fmt.Errorf("There are no 2 values, can't calculate expression - %s", char)
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
			//fmt.Printf(char)
			var err error
			if value, err = strconv.ParseFloat(char, 64); err != nil {
				return "Nil", fmt.Errorf("Wrong value - \"%s\"", char)
			}
		}
		stack = append(stack, value)
	}
	if len(stack) != 1 {
		return "Nil", fmt.Errorf("Wrong expression")
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
