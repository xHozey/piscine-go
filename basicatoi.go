package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, num := range s {
		num1 := int(num - '0')
		result = result*10 + num1

	}
	return result
}
