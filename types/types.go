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

	// Arrays
	var a1 = [2]string{"hey", "hi"}
	fmt.Println(a1[0], a1[1])
	fmt.Println(len(a1))
}
