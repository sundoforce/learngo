package main

import "fmt"

type pserson struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := pserson{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)

}
