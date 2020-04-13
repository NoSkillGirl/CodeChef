package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		count := 0
		for num > 0 {
			x := num % 10
			num = num / 10
			if x == 4 {
				count = count + 1
			}
		}
		fmt.Println(count)
	}
}
