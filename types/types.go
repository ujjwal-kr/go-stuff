package main

import (
	"fmt"
)

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

	// Maps
	m := make(map[string]string)
	p := make(map[string]int)
	m["name"] = "Ujjwal"
	m["role"] = "dev"
	p["age"] = 16
	fmt.Println(m["name"], m["role"], p["age"])

	// Pointers

	v := 5
	pv := &v
	fmt.Println(v, pv)
	*pv = 4
	fmt.Println(v, pv)

	// Two pointers of same value have the same address

	pv2 := &v
	fmt.Println(pv2)

	// Structs and types

	type Request struct {
		Resource string
	}
	type AuthenticatedRequest struct {
		Request
		Username, Password string
	}
	ar := new(AuthenticatedRequest)
	ar.Request.Resource = "https://google.com"
	ar.Resource = "https://google.com" // Same as above
	ar.Username = "Ujjwal"
	ar.Password = "123"
	fmt.Println(ar.Resource, ar.Username)
}
