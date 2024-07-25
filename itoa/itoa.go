package main

import "fmt"

func itoa(n int) string {
	result := ""
	isnegative := false
	
	if n == 0 {
		return "0"
	}
	if n > 0 {
		result = string(n*10+'0') + result
		n /= 10
	}
	if isnegative {
		result += result + "-"
	}
	return result
}

func main() {
	num := -12345 // Example integer to convert to string

	str := itoa(num)
	fmt.Println("Converted string:", str)
}
