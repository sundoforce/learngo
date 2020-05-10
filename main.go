package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 5
	fmt.Println(*b)

	a = 2
	b = &a
	*b = 20
	fmt.Println(a)
	fmt.Println(&b)
}
