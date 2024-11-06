package main

import "fmt"

func minnn(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n int
	var b string
	var m int
	var u bool
	var pr string
	var s []string
	fmt.Println("Сколько строк")
	fmt.Scan(&n)
	fmt.Println("Вводите строки")
	m = 1000000000000000
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		m = minnn(m, len(b))
		s = append(s, b)
	}
	for i := 0; i < m; i++ {
		u = true
		for j := 0; j < len(s)-1; j++ {
			if s[j][i] != s[j+1][i] {
				u = false
			}
		}
		if u == true {
			pr += string(rune(s[0][i]))
		} else {
			break
		}
	}
	fmt.Println(pr)

}
