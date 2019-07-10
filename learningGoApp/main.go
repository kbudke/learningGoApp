package main

import (
	"fmt"
	"os"
	"strconv"

	a "./area"
	"./calculator"
)

func main() {

	//getOperation is a label, use labels to define sections within a single func or multi func
	//we use as a checkpoint under main, to choose an operation.
getOperation:
	var operation string
	fmt.Print("Do you want Arithmetic(+, -, *, /) or Geometry(Area of shapes)?")
	fmt.Scanln(&operation)

	switch *&operation {
	case "Arithmetic", "arithmetic", "+", "-", "*", "/":
		fmt.Println("Arithmetic it is!")
		fmt.Scanln(&operation)

	case "Geometry", "geometry", "Area", "area":
		fmt.Println("Area it is!")
		fmt.Scanln(&operation)

		var arithmetic string
		fmt.Print("Please select an operation: +, -, *, / : ")
		fmt.Scanln(&arithmetic)
		var num1 string
		fmt.Print("Input the first number: ")
		fmt.Scanln(&num1)

		var num2 string
		fmt.Print("Input the second number: ")
		fmt.Scanln(&num2)

		// var mathOperation string
		// var operation string

		switch arithmetic {
		case "+":
			var firstNumber = stringToInt(num1)
			var secondNumber = stringToInt(num2)
			fmt.Print("Result: ")
			fmt.Println(calculator.Add(firstNumber, secondNumber))
		case "-":
			var firstNumber = stringToInt(num1)
			var secondNumber = stringToInt(num2)
			fmt.Print("Result: ")
			fmt.Println(calculator.Subtract(firstNumber, secondNumber))
		case "*":
			var firstNumber = stringToFloat64(num1)
			var secondNumber = stringToFloat64(num2)
			fmt.Print("Result: ")
			fmt.Println(calculator.Multiply(firstNumber, secondNumber))
		case "/":
			var firstNumber = stringToFloat64(num1)
			var secondNumber = stringToFloat64(num2)
			fmt.Print("Result: ")
			fmt.Println(calculator.Divide(firstNumber, secondNumber))

		default:
			fmt.Println("Invalid operation selected. Please try again!")
			//The goto in this case saves us from introducing another (boolean) variable used just for control-flow, checked for at the end. In this case, the goto statement makes the code actually better to read and easier follow
			goto getOperation
		}

		var geometry string
		switch geometry {
		case "circle":
			var radius float64
			var circle a.Circle
			var area = stringToFloat64(circle)
			fmt.Print("Result: ")
			fmt.Println(a.Circle(circle.radius))
		case "square":
			var square a.Square
			var area = stringToFloat64(square)
			fmt.Print("Result: ")
			fmt.Println(a.Square(a.Square.length))

			fmt.Print("Select the shape you want area Calculated: Circle, Square : ")
			fmt.Scanln(&calculatorOperation)
			var circle float64
			fmt.Print("Input the Radius of the circle: ")
			fmt.Scanln(&radius)
			var square float64
			fmt.Print("Input a Length of the square: ")
			fmt.Scanln(&length)

		default:
			fmt.Println("Invalid operation selected. Please try again!")
			//The goto in this case saves us from introducing another (boolean) variable used just for control-flow, checked for at the end. In this case, the goto statement makes the code actually better to read and easier follow
			goto getOperation
		}
	}
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
