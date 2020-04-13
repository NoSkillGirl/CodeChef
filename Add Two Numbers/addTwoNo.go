package main

import "fmt"

func main() {
	var (
		a int
		b int
		n int
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
