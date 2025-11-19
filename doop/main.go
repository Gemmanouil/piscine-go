package main

import "os"

func Atoi(s string) (int, bool) {
	sign := 1
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	if s == "" {
		return 0, false
	}
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		n = n*10 + int(r-'0')
	}
	return n * sign, true
}

func PrintNbr(n int) {
	if n == 0 {
		PrintStr("0\n")
		return
	}
	str := ""
	if n < 0 {
		str += "-"
		n = -n
	}
	digits := ""
	for n > 0 {
		digits = string(rune(n%10+'0')) + digits
		n /= 10
	}
	PrintStr(str + digits + "\n")
}

func PrintStr(s string) {
	for _, r := range s {
		os.Stdout.Write([]byte{byte(r)})
	}
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := Atoi(os.Args[1])
	b, ok2 := Atoi(os.Args[3])
	op := os.Args[2]

	if !ok1 || !ok2 {
		return
	}

	if op == "/" && b == 0 {
		PrintStr("No division by 0\n")
		return
	}
	if op == "%" && b == 0 {
		PrintStr("No modulo by 0\n")
		return
	}

	if op == "+" {
		PrintNbr(a + b)
	} else if op == "-" {
		PrintNbr(a - b)
	} else if op == "*" {
		PrintNbr(a * b)
	} else if op == "/" {
		PrintNbr(a / b)
	} else if op == "%" {
		PrintNbr(a % b)
	}
}
