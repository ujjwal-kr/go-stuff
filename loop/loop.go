package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while

	i := 0
	for i < 100 {
		fmt.Println(i)
		i += 10
	}
}
