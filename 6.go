package main

import "fmt"
import "math"

func main() {
	var a, b int
	var o string
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&o)

	switch o {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("деление на ноль")
		}
		fmt.Println(a / b)
	case "^":
		fmt.Println(int(math.Pow(float64(a), float64(b))))
	case "%":
		fmt.Println(a % b)
	default:
		fmt.Println("недопустимая операция")
	}
}
