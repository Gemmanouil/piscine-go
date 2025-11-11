package piscine

// IterativeFactorial υπολογίζει το παραγοντικό ενός αριθμού nb
// Αν το nb είναι αρνητικό ή πολύ μεγάλο (overflow), επιστρέφει 0
func IterativeFactorial(nb int) int {
	// Έλεγχος για μη επιτρεπτές τιμές
	if nb < 0 {
		return 0 // αρνητικοί αριθμοί δεν έχουν παραγοντικό
	}

	result := 1 // ξεκινάμε με 1 γιατί το παραγοντικό πολλαπλασιάζει

	// βρόχος από το 1 μέχρι το nb
	for i := 1; i <= nb; i++ {
		result *= i // πολλαπλασιάζουμε το result με το i

		// απλός έλεγχος overflow: αν το αποτέλεσμα γίνει αρνητικό (ξεχείλισμα int)
		if result < 0 {
			return 0
		}
	}

	return result // επιστρέφουμε το αποτέλεσμα
}
