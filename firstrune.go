package piscine

// FirstRune επιστρέφει τον πρώτο χαρακτήρα (rune) της συμβολοσειράς s
func FirstRune(s string) rune {
	for _, r := range s {
		return r // επιστρέφουμε τον πρώτο χαρακτήρα που βρίσκουμε
	}
	return 0 // αν η συμβολοσειρά είναι κενή, επιστρέφουμε 0
}
