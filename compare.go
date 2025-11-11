package piscine

// Compare συγκρίνει δύο συμβολοσειρές και επιστρέφει:
// 0 αν είναι ίδιες
// αρνητικό αν a < b
// θετικό αν a > b
func Compare(a, b string) int {
	// Μετατρέπουμε τις συμβολοσειρές σε slices από runes για σωστή σύγκριση Unicode // (listes * python)
	ar := []rune(a)
	br := []rune(b)

	// Βρίσκουμε το μικρότερο μήκος για να συγκρίνουμε χαρακτήρα-χαρακτήρα
	minLen := len(ar)
	if len(br) < minLen {
		minLen = len(br)
	}

	// Συγκρίνουμε χαρακτήρα-χαρακτήρα
	for i := 0; i < minLen; i++ {
		if ar[i] != br[i] {
			return int(ar[i]) - int(br[i]) // επιστρέφουμε τη διαφορά των χαρακτήρων
		}
	}

	// Αν όλοι οι χαρακτήρες είναι ίδιοι, συγκρίνουμε τα μήκη
	return len(ar) - len(br)
}
