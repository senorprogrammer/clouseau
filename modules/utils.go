package modules

import ()

func contains(source []string, target string) bool {
	for _, el := range source {
		if el == target {
			return true
		}
	}
	return false
}

func stripQuotes(source string) string {
	quotes := []byte{'\'', '"'}

	for _, quote := range quotes {
		if len(source) > 0 && source[0] == quote {
			source = source[1:]
		}
		if len(source) > 0 && source[len(source)-1] == quote {
			source = source[:len(source)-1]
		}
	}

	return source
}
