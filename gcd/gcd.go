package main

import (
	"fmt"
	//"piscine"
)

func Gcd(a, b int) int {
	// Keep applying the Euclidean algorithm until b becomes 0
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
