package piscine

// Compare συγκρίνει δύο συμβολοσειρές και επιστρέφει:
// 0 αν είναι ίδιες
// -1 αν a < b
// 1 αν a > b
func Compare(a, b string) int {
	ar := []rune(a) // sumboloseires (listes *python)
	br := []rune(b)

	minLen := len(ar)
	if len(br) < minLen {
		minLen = len(br)
	}

	for i := 0; i < minLen; i++ {
		if ar[i] < br[i] {
			return -1
		}
		if ar[i] > br[i] {
			return 1
		}
	}

	if len(ar) < len(br) {
		return -1
	}
	if len(ar) > len(br) {
		return 1
	}

	return 0
}
