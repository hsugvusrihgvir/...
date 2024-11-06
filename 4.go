package main

import (
	"fmt"
)

var m1 = [...]int{1, 3, 5, 7, 9}
var m2 = [...]int{2, 4, 6, 8, 10, 11, 12, 13}
var m [len(m1) + len(m2)]int
var o1, o2 int = 0, 0

func main() {
	for i := 0; i < len(m1)+len(m2); i++ {
		if o1 < len(m1) && (o2 >= len(m2) || m1[o1] <= m2[o2]) {
			m[i] = m1[o1]
			o1++
		} else {
			m[i] = m2[o2]
			o2++
		}
	}
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}
