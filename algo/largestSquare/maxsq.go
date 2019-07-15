package main

import (
	"fmt"
)

func sqSize(left, right int) int {
	if left == right {
		return right
	}
	first := left
	second := right
	if right > left {
		first = right
		second = left
	}
	for first > second {
		first = first - second
	}
	return sqSize(second, first)
}

func main() {
	fmt.Print("Enter Plot Size (2 Integers): ")
	var first, second int
	fmt.Scan(&first, &second)
	fmt.Println("Plot: ", first, " by ", second)
	fmt.Println("Smallest Square: ", sqSize(first, second))
}
