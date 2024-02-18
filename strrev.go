package piscine

func StrRev(s string) string {
	runex := []rune(s)
	for x, h := 0, len(runex)-1; x < h; x, h = x+1, h-1 {
		runex[x], runex[h] = runex[h], runex[x]
	}

	return string(runex)
}
