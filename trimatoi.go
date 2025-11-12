package piscine

func TrimAtoi(s string) int {
	// This variable will hold the sign of the number (positive or negative)
	sign := 1

	// This flag tells us if we've started collecting digits or seen a sign
	started := false

	// This is the final result we will build from digits
	result := 0

	// Loop through each character in the string
	for _, char := range s {
		// If we find a '-' and haven't started yet, set the sign to negative
		if char == '-' && !started {
			sign = -1
			started = true
		} else if char >= '0' && char <= '9' {
			// If the character is a digit, mark that we've started
			started = true

			// Build the number by shifting digits left and adding the new one
			result = result*10 + int(char-'0')
		}
		// Ignore all other characters
	}

	// If no digits were found, return 0
	if !started {
		return 0
	}

	// Apply the sign and return the final number
	return result * sign
}
