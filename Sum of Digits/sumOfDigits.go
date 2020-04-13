package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for n > 0 {
		var num int
		fmt.Scan(&num)
		r := 0
		for {
			if num == 0 {
				break
			} else {
				r = r + num%10
				num = num / 10
			}
		}
		fmt.Println(r)
		n--
	}
}

