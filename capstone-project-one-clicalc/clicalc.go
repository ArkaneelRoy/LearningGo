package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}

func remainder(a, b float64) (int64, bool) {
	if int64(b) == 0 {
		return 0, false
	} else {
		return int64(a) % int64(b), true
	}
}

func product(a, b float64) float64 {
	return a * b
}

func power(a float64, b float64) float64 {
	res := 1.0
	for i := 0; i < int(b); i++ {
		res *= a
	}
	return res
}

func main() {
	for {
		// Parsing Logic

		var num1, num2 float64
		var operator string
		fmt.Print("Enter your first number :")
		fmt.Scanln(&num1)
		fmt.Print("Enter any of the operator (+,-,*,/,%,^) :")
		fmt.Scanln(&operator)
		fmt.Print("Enter your second number :")
		fmt.Scanln(&num2)

		// operation logic

		fmt.Println("The operation is :")
		switch operator {
		case "+":
			fmt.Println(add(num1, num2))
		case "-":
			fmt.Println(sub(num1, num2))
		case "/":
			res, ok := divide(num1, num2)
			if !ok {
				fmt.Println("Undefined...")
			} else {
				fmt.Println(res)
			}
		case "*":
			fmt.Println(product(num1, num2))
		case "%":
			res, ok := remainder(num1, num2)
			if !ok {
				fmt.Println("Undefined...")
			} else {
				fmt.Println(res)
			}
		case "^":
			fmt.Println(power(num1, num2))
		default:
			fmt.Println("Unsupported...")
		}

		// exit logic

		var choice string
		fmt.Print("Please enter 'exit' to exit the program or 'continue' to continue the program:")
		fmt.Scanln(&choice)
		if choice == "exit" {
			fmt.Println("Youre Leaving me (╥﹏╥)")
			break
		} else if choice == "continue" {
			fmt.Println("Continuing Calculation...")
		}
	}
}
