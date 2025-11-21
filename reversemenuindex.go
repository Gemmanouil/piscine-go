package piscine

func ReverseMenuIndex(menu []string) []string {
	// Create a new slice with the same length as the input
	reversedMenu := make([]string, len(menu))

	// Fill the new slice with elements from the original in reverse order
	for i := 0; i < len(menu); i++ {
		// reversedMenu[0] gets menu[len(menu)-1], and so on
		reversedMenu[i] = menu[len(menu)-1-i]
	}

	return reversedMenu
}
