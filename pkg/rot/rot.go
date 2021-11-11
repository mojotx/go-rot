package rot

// Rot is a rot13 function in Go
func Rot(r rune) rune {

	if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
		return r
	}

	if r >= 'A' && r <= 'M' || r >= 'a' && r <= 'm' {
		return r + 13
	} else {
		return r - 13
	}
}
