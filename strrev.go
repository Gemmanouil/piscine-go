package piscine

func StrRev(s string) string {
	runes := []rune(s)
	lenght := len(runes)

	for i := 0; i < lenght/2; i++ {
		runes[i], runes[lenght-1-i] = runes[lenght-1-i], runes[i]
	}

	return string(runes)
}
