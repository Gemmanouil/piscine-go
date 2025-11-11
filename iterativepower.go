package piscine

// IterativePower υπολογίζει το nb^power χρησιμοποιώντας επανάληψη (όχι αναδρομή)
func IterativePower(nb int, power int) int {
	// Αν η δύναμη είναι αρνητική, επιστρέφουμε 0 σύμφωνα με τις οδηγίες
	if power < 0 {
		return 0
	}

	// Ξεκινάμε με αποτέλεσμα = 1 γιατί κάθε αριθμός υψωμένος στο 0 είναι 1
	result := 1

	// Επαναλαμβάνουμε power φορές, πολλαπλασιάζοντας το αποτέλεσμα με το nb
	for i := 0; i < power; i++ {
		result *= nb
	}

	// Επιστρέφουμε το τελικό αποτέλεσμα
	return result
}
