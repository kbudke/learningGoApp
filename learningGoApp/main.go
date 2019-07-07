package main

import (
	"fmt"
	"os"
	"strconv"

	tryHSTS "./http"
	a "./area"
	"./calculator"
)

func main() {

	hsts := tryHSTS.serveHTTPS() {

	}
	
	func hsts() {
	if *serveHTTPSFlag {
		serveHTTPS()
	} else {
		serveHTTP()
	}
	}
	//TODO: finish area calculations

	// circ := a.Circle{
	// 	radius: 12.345,
	// }
	// sqr := a.Square{
	// 	length: 15,
	// }
	// info(circ)
	// info(sqr)


	// TODO: set up the first switch for is this a regular math opperation
	// TODO: or is this to calculate area?

	hours, remainder := Divmod(5566, 3600)
	minutes, seconds := Divmod(remainder, 60)
	fmt.Printf("%d:%d:%d", hours, minutes, seconds)

	//getOperation is a label, use labels to define sections within a single func or multi func
	//we use as a checkpoint under main, to choose an operation.
getOperation:
	var operation string
	fmt.Print("Select an operation: +, -, *, / : ")
	//Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
	fmt.Scanln(&operation)

	if area {
		fmt.Print("Area can help with Calculating the area of the following shapes.")
		fmt.Println(area.Square)
	}

	var num1 string
	fmt.Print("Input the first number: ")
	fmt.Scanln(&num1)

	var num2 string
	fmt.Print("Input the second number: ")
	fmt.Scanln(&num2)

	switch operation {
	case "+":
		var firstNumber int = stringToInt(num1)
		var secondNumber int = stringToInt(num2)
		fmt.Print("Result: ")
		fmt.Println(calculator.Add(firstNumber, secondNumber))
	case "-":
		var firstNumber int = stringToInt(num1)
		var secondNumber int = stringToInt(num2)
		fmt.Print("Result: ")
		fmt.Println(calculator.Subtract(firstNumber, secondNumber))

	case "*":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(calculator.Multiply(firstNumber, secondNumber))

	case "/":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(calculator.Divide(firstNumber, secondNumber))

	case "circle":
		var radius float64
		fmt.Print("Result: ")
		fmt.Println(a.Circle(radius))

	case "square":
		var length float64

		fmt.Print("Result: ")
		fmt.Println(a.Square(length))

	default:
		fmt.Println("Invalid operation selected. Please try again!")
		//The goto in this case saves us from introducing another (boolean) variable used just for control-flow, checked for at the end. In this case, the goto statement makes the code actually better to read and easier follow
		goto getOperation
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
