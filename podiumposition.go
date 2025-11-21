package piscine

func PodiumPosition(podium [][]string) [][]string {
	// Create a new slice with the same length as the input
	reversed := make([][]string, len(podium))

	// Fill the new slice in reverse order
	for i := 0; i < len(podium); i++ {
		reversed[i] = podium[len(podium)-1-i]
	}

	return reversed
}
