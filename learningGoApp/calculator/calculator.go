package calculator

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Add(num1, num2 int) int {
	return num1 + num2
}

//Subtract is
func Subtract(num1, num2 int) int {
	return num1 - num2
}

//Multiply is
func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

//Divide is
func Divide(num1, num2 float64) float64 {
	return num1 / num2
}

//Divmod is
func Divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

//Sqrt is
func Sqrt(num float64) float64 {
	return math.Sqrt(num)
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func stringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return f
}

func float64ToString(float float64) string {
	return fmt.Sprintf("%f", float)
}
func calculateExpression(expr []string) float64 {
	var solution float64
	// remainingExpr := append(expr)
evaluateMore:
	for i := 0; i < len(expr); i++ {
		switch expr[i] {
		case "+":
			var calculation = stringToFloat64(expr[i-1]) + stringToFloat64(expr[i+1])
			ans := []string{float64ToString(calculation)}
			expr = append(ans, expr[i+2:]...)
			goto evaluateMore
		case "-":

		case "*":

		case "=":
			solution = stringToFloat64(expr[i-1])
		}

	}
	return solution
}

//ArithmeticCalculator main controller
func ArithmeticCalculator(arithmetic string) {
	executionArray := []string{}
getOperation:
	// var arithmetic string
	// fmt.Print("Please select an operation: +, -, *, / : ")
	// fmt.Scanln(&arithmetic)

	for {
		// TODO: can we note the thought process of anything random set up and or why we went with it..
		var input string
		fmt.Scanln(&input)
		if input != "=" {
			executionArray = append(executionArray, input)
		} else {
			executionArray = append(executionArray, input)
			solution := calculateExpression(executionArray)
			fmt.Println(strings.Join(executionArray, " "), solution)
		}
		// fmt.Println(executionArray)
		// var num1 string
		// fmt.Print("Input the first number: ")
		// fmt.Scanln(&num1)

		// var num2 string
		// fmt.Print("Input the second number: ")
		// fmt.Scanln(&num2)

		// switch input {
		// case "+":
		// 	var firstNumber = stringToInt(num1)
		// 	var secondNumber = stringToInt(num2)
		// 	fmt.Print("Result: ")
		// 	fmt.Println(Add(firstNumber, secondNumber))
		// case "-":
		// 	var firstNumber = stringToInt(num1)
		// 	var secondNumber = stringToInt(num2)
		// 	fmt.Print("Result: ")
		// 	fmt.Println(Subtract(firstNumber, secondNumber))
		// case "*":
		// 	var firstNumber = stringToFloat64(num1)
		// 	var secondNumber = stringToFloat64(num2)
		// 	fmt.Print("Result: ")
		// 	fmt.Println(Multiply(firstNumber, secondNumber))
		// case "/":
		// 	var firstNumber = stringToFloat64(num1)
		// 	var secondNumber = stringToFloat64(num2)
		// 	fmt.Print("Result: ")
		// 	fmt.Println(Divide(firstNumber, secondNumber))
		// // case "back":
		// // 	return
		// case "=":

		// default:
		// 	fmt.Println("Invalid operation selected. Please try again!")
		// 	//The goto in this case saves us from introducing another (boolean) variable used just for control-flow, checked for at the end. In this case, the goto statement makes the code actually better to read and easier follow
		// 	goto getOperation
		// }

		// fmt.Print("Please select an operation: +, -, *, / : ")
		// fmt.Scanln(&arithmetic)
		goto getOperation
	}

}
