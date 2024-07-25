package main

func atoi(s string) int {
	var result int
	 sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		sign = 1
		s = s[1:]
	}
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = (result * 10) + int(char-'0')
		} else {
			return 0
		}
	}
	result *= sign
	return result
}
