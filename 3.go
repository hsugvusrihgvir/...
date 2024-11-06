package main

import "fmt"

func main() {
	var n int
	var l = []int{1, 1}
	var ll = []int{1}
	fmt.Scan(&n)
	fmt.Println(1)
	for i := 0; i < n-1; i++ {
		for j := 0; j < len(l); j++ {
			fmt.Print(l[j], " ")
		}
		fmt.Println(" ")
		for j := 1; j < len(l); j++ {
			ll = append(ll, l[j-1]+l[j])
		}
		ll = append(ll, 1)
		l = ll
		ll = []int{1}

	}
}
