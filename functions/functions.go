package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(20, 30))
}

func add(a int, b int) int {
	sum := a + b
	return sum
}
