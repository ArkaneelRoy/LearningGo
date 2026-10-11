package main

func bitAND(a, b int64) int64 {
	return a & b
}

func bitOR(a, b int64) int64 {
	return a | b
}

func bitXOR(a, b int64) int64 {
	return a ^ b
}

func shiftLeft(a, b int64) int64 {
	return a << uint(b)
}

func shiftRight(a, b int64) int64 {
	return a >> uint(b)
}

/** func main() {
	// Below these are only for testing
	fmt.Printf("3 is binary : %04b\n", 3)
	fmt.Printf("4 is binary : %04b\n", 4)
	fmt.Printf("ResultAND : %4b (Decimal : %d)\n", bitAND(3, 4), bitAND(3, 4))
	fmt.Printf("ResultOR : %4b (Decimal : %d)\n", bitOR(3, 4), bitOR(3, 4))
	fmt.Printf("ResultXOR : %4b (Decimal : %d)\n", bitXOR(3, 4), bitXOR(3, 4))
	fmt.Printf("ResultLeft : %4b (Decimal : %d)\n", shiftLeft(3, 4), shiftLeft(3, 4))
	fmt.Printf("ResultRight : %4b (Decimal : %d)\n", shiftRight(3, 4), shiftRight(3, 4))
} **/
