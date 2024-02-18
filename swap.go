package piscine

func Swap(a *int, b *int) {
	stock := *b
	*b = *a
	*a = stock
}
