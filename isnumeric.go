package piscine

func IsNumeric(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return false // αν βρούμε μη αλφαριθμητικό χαρακτήρα, επιστρέφουμε false
		}
	}
	return true
}
