//
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n > 0 {
		var a int
		fmt.Scan(&a)
		fmt.Println(fact(a))
		n--
	}
}

func fact(a int) int {
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	return a * fact(a-1)

}
