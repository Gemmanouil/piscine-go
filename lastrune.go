package piscine

// LastRune επιστρέφει τον τελευταίο χαρακτήρα (rune) της συμβολοσειράς s
func LastRune(s string) rune {
	var last rune // μεταβλητή για να κρατήσουμε τον τελευταίο χαρακτήρα
	for _, r := range s {
		last = r // κάθε φορά κρατάμε τον επόμενο χαρακτήρα
	}
	return last // στο τέλος, αυτός θα είναι ο τελευταίος
}
