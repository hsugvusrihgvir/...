package main

import "fmt"

var g int

func main() {
	fmt.Scan(&g)
	if (g%4 == 0 && g%100 != 0) || (g%400 == 0) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}
}
