package piscine

func ConcatParams(args []string) string {
	var conCatStrings string
	for j, i := range args {
		if j < len(args) {

			conCatStrings += i
			conCatStrings += "\n"
		} else {
			conCatStrings += i
		}
	}

	return conCatStrings
}
