package calculator

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

//MATH Functions
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

// HELPER Functions
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
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

func splitStringByCharacter(str string) []string {
	strArray := []string{}
	var multiCharValue string
	index := 0
	for _, char := range str {
		if isNumeric(string(char)) && ((index + 1) != len(str)) && isNumeric(string(str[index+1])) {
			multiCharValue += string(char)
		} else if isNumeric(string(char)) && ((index + 1) != len(str)) && !isNumeric(string(str[index+1])) {
			multiCharValue += string(char)
			strArray = append(strArray, string(multiCharValue))
			multiCharValue = ""
		} else {
			strArray = append(strArray, string(char))
		}
		index++
	}
	return strArray
}

func contains(a []string, x string) int {
	var index int
	for i, n := range a {
		if x == n {
			index = i
			break //not greedy will return after first occurence found
		} else {
			index = -1
		}

	}
	return index
}

func evaluateExpression(expr []string) []string {
	for {
		fmt.Println("evaluate expr ", expr)
		if len(expr) >= 3 {
			subExpr := expr[0:3]
			subSolution := evaluate(subExpr)
			subSolutionStr := float64ToString(subSolution)
			//update the expr array
			subExprSolution := []string{subSolutionStr}
			suffix := expr[3:]

			expr = append(subExprSolution, suffix...)
		} else {
			break
		}
	}
	return expr
}

func evaluate(expr []string) float64 {
	var solution float64
	switch expr[1] {
	case "+":
		solution = stringToFloat64(expr[0]) + stringToFloat64(expr[2])
	case "-":
		solution = stringToFloat64(expr[0]) - stringToFloat64(expr[2])
	case "*":
		solution = stringToFloat64(expr[0]) * stringToFloat64(expr[2])
	}
	return solution
}

func calculateExpression(expr []string) float64 {
	var solution float64

	for {
		//search for operators following PEMDAS precedences
		if openParenthesisIndex := contains(expr, "("); openParenthesisIndex != -1 {
			closeParentheisIndex := contains(expr, ")")
			subExpr := expr[(openParenthesisIndex + 1):(closeParentheisIndex)]
			subExpr = evaluateExpression(subExpr)
			//reset the expr array; replace (expr) with solution
			prefixAndSubExprSolution := append(expr[0:openParenthesisIndex], subExpr[0])
			suffix := expr[(closeParentheisIndex + 1):]
			expr = append(prefixAndSubExprSolution, suffix...)
		} else { //addition or subtraction
			expr = evaluateExpression(expr)
		}
		if len(expr) == 1 {
			solution = stringToFloat64(expr[0])
			break
		}
	}
	return solution
}

//ArithmeticCalculator main controller
func ArithmeticCalculator(arithmetic string) {
	var input string

	for {
		fmt.Println("Enter the expression on a single line. NO SPACES. NO EQUALS SIGN")
		fmt.Println("Available inputs: [ 0-9 ( ) + - * ]")
		// TODO: can we note the thought process of anything random set up and or why we went with it..
		fmt.Scanln(&input)
		solution := calculateExpression(splitStringByCharacter(input))
		fmt.Println()
		fmt.Println(input, "= ", solution)
		fmt.Println()

	}

}
