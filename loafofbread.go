package piscine

func LoafOfBread(s string) string {
	// Count non-space characters
	count := 0
	for _, r := range s {
		if r != ' ' {
			count++
		}
	}

	// Special case: empty or only spaces → just newline
	if count == 0 {
		return "\n"
	}

	// Less than 5 non-space chars → Invalid Output
	if count < 5 {
		return "Invalid Output\n"
	}

	result := ""
	temp := ""
	skipNext := false

	for _, r := range s {
		if skipNext {
			skipNext = false
			continue
		}

		if r == ' ' {
			continue
		}

		temp += string(r)

		// When we have a full 5-char block
		if len(temp) == 5 {
			if result != "" {
				result += " "
			}
			result += temp
			temp = ""
			skipNext = true
		}
	}

	// Add leftover characters (even if <5)
	if temp != "" {
		result += " " + temp
	}

	return result + "\n"
}
