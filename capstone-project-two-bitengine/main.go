package main

import "fmt"

func main() {
	for {
		var num1, num2 int64
		var operator string
		fmt.Print("Enter the first int64 :")
		_, err := fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("Error reading input :", err)
			return
		}

		fmt.Print("Enter operator (&, |, ^, <<, >>, divmod, pow) or 'exit' : ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("Shutting down engine..")
			break
		}

		fmt.Print("Enter the second int64 :")
		fmt.Scanln(&num2)

		switch operator {
		case "&":
			res := bitAND(num1, num2)
			fmt.Printf("Decimal: %d | Binary: %016b\n", res, uint16(res))
		case "|":
			res := bitOR(num1, num2)
			fmt.Printf("Decimal: %d | Binary: %016b\n", res, uint16(res))
		case "^":
			res := bitXOR(num1, num2)
			fmt.Printf("Decimal: %d | Binary: %016b\n", res, uint16(res))
		case "<<":
			if num2 < 0 {
				fmt.Println("Error: Shift count cannot be negative")
			} else {
				res := shiftLeft(num1, num2)
				fmt.Printf("Decimal: %d | Binary: %016b\n", res, uint16(res))
			}
		case ">>":
			if num2 < 0 {
				fmt.Println("Error: Shift count cannot be negative")
			} else {
				res := shiftRight(num1, num2)
				fmt.Printf("Decimal: %d | Binary: %016b\n", res, uint16(res))
			}
		case "divmod":
			q, r, ok := divMod(num1, num2)
			if !ok {
				fmt.Println("Hardware Fault: Division by zero")
			} else {
				fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
			}
		case "pow":
			res, ok := safePow(num1, num2)
			if !ok {
				fmt.Println("Hardware Fault: Integer overflow or negative exponent")
			} else {
				fmt.Println("Result:", res)
			}
		default:
			fmt.Println("Unsupported operation")
		}
    	fmt.Println("----")
	}
}
