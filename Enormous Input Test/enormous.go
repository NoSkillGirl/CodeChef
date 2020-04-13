package main

import "fmt"

func main() {
	var (
		a int
		b int
		te int
	)
	fmt.Scan(&a, &b)
	cn := 0
	for i := 0; i < a; i++ {
		// var te int
		fmt.Scan(&te)
		if te%b == 0 {
			cn ++
		}
	}
	fmt.Println(cn)
}
