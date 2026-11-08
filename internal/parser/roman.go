package parser

func isRomanNumeral(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if romanValue(r) == 0 {
			return false
		}
	}
	if romanToInt(s) <= 0 {
		return false
	}
	return true
}

func romanToInt(s string) int {
	total, previous := 0, 0
	for _, r := range s {
		value := romanValue(r)
		if value > previous {
			total += value - 2*previous
		} else {
			total += value
		}
		previous = value
	}
	return total
}

func romanValue(r rune) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
