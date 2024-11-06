package main

import "fmt"

var s string
var ps string
var f bool
var k int
var ind int

func main() {
	fmt.Scan(&s)
	fmt.Scan(&ps)
	for i := 0; i < len(s); i++ {
		if f == false {
			if s[i] == ps[0] {
				f = true
				k = 1
				ind = i
			}
		} else {
			if s[i] == ps[k] {
				k += 1
				if k >= len(ps) {
					break
				}
			} else {
				k = 0
				f = false
			}
		}
	}
	if f == true {
		fmt.Println(ind)
	} else {
		fmt.Println("нет")
	}
}
