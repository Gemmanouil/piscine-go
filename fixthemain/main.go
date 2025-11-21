package piscine

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	OPEN  = true
	CLOSE = false
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("Is the Door opened?")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed?")
	return ptrDoor.state == CLOSE
}
