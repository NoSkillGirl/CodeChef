package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var first int
		fmt.Scan(&num)
		last := num % 10
		for num > 0 {
			first = num % 10
			num = num / 10
		}

		fmt.Println(first + last)
	}
}
