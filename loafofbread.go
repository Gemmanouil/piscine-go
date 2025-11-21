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
	collected := 0
	skipNext := false

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skipNext {
			skipNext = false
			continue
		}

		result = append(result, r)
		collected++

		if collected == 5 {
			result = append(result, ' ')
			collected = 0
			skipNext = true
		}
	}

	// Remove trailing space if present
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	result = append(result, '\n')

	return string(result)
}
