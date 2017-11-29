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

func removeDuplicates(source []string) []string {
	exists := map[string]bool{}
	result := []string{}

	for v := range source {
		if exists[source[v]] != true {
			exists[source[v]] = true
			result = append(result, source[v])
		}
	}

	return result
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
