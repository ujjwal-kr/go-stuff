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

	// loop over array
	a := []int{1, 2, 3}
	fmt.Println(a)

	for idx, value := range a {
		fmt.Println(idx, value)
	}
}
