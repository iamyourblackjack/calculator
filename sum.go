package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers
func Sum(num1, num2 int) int {
	return num1 + num2
}