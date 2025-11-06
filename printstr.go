// str
package piscine

func PrintStr(s string) {
	for range s {
		PrintStr(s)
	}
}
