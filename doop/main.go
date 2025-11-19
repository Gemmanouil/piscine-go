package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check for exactly 3 arguments (excluding program name)
	if len(os.Args) != 4 {
		return // Invalid number of arguments
	}

	// Parse the first and third arguments as integers
	a, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	b, err2 := strconv.ParseInt(os.Args[3], 10, 64)
	op := os.Args[2]

	// If either value is not a valid integer, exit silently
	if err1 != nil || err2 != nil {
		return
	}

	// Handle division or modulo by zero
	if (op == "/" || op == "%") && b == 0 {
		if op == "/" {
			fmt.Println("No division by 0")
		} else {
			fmt.Println("No modulo by 0")
		}
		return
	}

	// Perform the operation
	switch op {
	case "+":
		result := a + b
		if (a > 0 && b > 0 && result < 0) || (a < 0 && b < 0 && result > 0) {
			return // Overflow
		}
		fmt.Println(result)
	case "-":
		result := a - b
		if (a > 0 && b < 0 && result < 0) || (a < 0 && b > 0 && result > 0) {
			return // Overflow
		}
		fmt.Println(result)
	case "*":
		result := a * b
		if b != 0 && result/b != a {
			return // Overflow
		}
		fmt.Println(result)
	case "/":
		fmt.Println(a / b)
	case "%":
		fmt.Println(a % b)
	default:
		return // Invalid operator
	}
}
