package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if (n%4 == 0 && n%100 != 0) || (n%400 == 0) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}
}
