package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n int
		a int
	)
	fmt.Scan(&n)
	sl := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sl[i] = a
	}
	sort.Ints(sl)
	for _, x := range sl {
		fmt.Println(x)
	}
}
