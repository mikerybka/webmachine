package util

func FilterNonDigits(s string) string {
	var filtered string
	for _, r := range s {
		if r >= '0' && r <= '9' {
			filtered += string(r)
		}
	}
	return filtered
}
