package piscine

func LoafOfBread(str string) string {
	// Step 1: Filter out spaces
	chars := []rune{}
	for _, r := range str {
		if r != ' ' {
			chars = append(chars, r)
		}
	}

	// Step 2: Check if we have at least 5 characters
	if len(chars) < 5 {
		return "Invalid Output\n"
	}

	// Step 3: Build result in chunks of 5, skipping the 6th
	result := []rune{}
	for i := 0; i < len(chars); {
		if len(chars)-i < 5 {
			break
		}
		// Add 5 characters
		for j := 0; j < 5; j++ {
			result = append(result, chars[i+j])
		}
		result = append(result, ' ') // add space after each 5-char chunk
		i += 6                       // skip the 6th character
	}

	// Step 4: Remove trailing space and add newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	result = append(result, '\n')

	return string(result)
}
