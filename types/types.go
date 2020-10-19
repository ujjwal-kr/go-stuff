package main

import "fmt"

func main() {
	// bool
	b := true
	fmt.Println(b)

	var a bool = true
	fmt.Println(a)

	// int
	var c int = 5
	d := 5
	fmt.Println(c + d)
	// convert to string
	str := fmt.Sprintf("%d", d)
	fmt.Println(str + " is good")

	const e string = "hey" // not re-assignable
	fmt.Println(e)
}
