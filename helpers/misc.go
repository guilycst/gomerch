package helpers

func Contains(s []string, t string) bool {
	for _, a := range s {
		if a == t {
			return true
		}
	}
	return false
}
