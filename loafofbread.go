package piscine

func LoafOfBread(str string) string {
	// Count non-space characters
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}
	if count < 5 {
		return "Invalid Output\n"
	}

	result := []rune{}
	word := []rune{}
	skip := false
	nonSpaceCount := 0

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skip {
			skip = false
			continue
		}
		word = append(word, r)
		nonSpaceCount++
		if nonSpaceCount%5 == 0 {
			result = append(result, word...)
			result = append(result, ' ')
			word = []rune{}
			skip = true
		}
	}

	// Remove trailing space and add newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	result = append(result, '\n')

	return string(result)
}
