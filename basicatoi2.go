package piscine

func BasicAtoi2(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		num := int(s[i]) - '0'
		result = result*10 + num
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
	}
	return result
}
