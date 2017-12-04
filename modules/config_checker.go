package modules

import (
	"regexp"
	"sort"
	"strings"
	"sync"
)

type ConfigChecker struct {
	name    string
	Results map[string][]string
	Path    string

	sync.Mutex
	regex *regexp.Regexp
}

func NewConfigChecker(name, path, searchStr string) *ConfigChecker {
	checker := ConfigChecker{
		name:    name,
		Results: make(map[string][]string),
		Path:    path,

		regex: regexp.MustCompile(searchStr),
	}

	return &checker
}

/* -------------------- Public Functions -------------------- */

func (checker *ConfigChecker) Keys() []string {
	keys := make([]string, 0, len(checker.Results))
	for key := range checker.Results {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (checker *ConfigChecker) Len() int {
	return len(checker.Results)
}

func (checker *ConfigChecker) Name() string {
	return checker.name
}

func (checker *ConfigChecker) Parse(data string, fileName string) {
	results := make(map[string][]string)

	for _, match := range checker.regex.FindAllString(data, -1) {
		results[match] = append(results[match], fileName)
	}

	checker.merge(results)
}

func (checker *ConfigChecker) Sanitize() {
	results := make(map[string][]string)

	for key, value := range checker.Results {
		key := checker.sanitizeKey(key)
		value := checker.sanitizePaths(value)
		value = removeDuplicates(value)
		results[key] = append(checker.Results[key], value...)
	}

	checker.Results = results
}

/* -------------------- Private Functions -------------------- */

func (checker *ConfigChecker) merge(results map[string][]string) {
	if len(results) == 0 {
		return
	}

	checker.Lock()
	defer checker.Unlock()

	for key, value := range results {
		checker.Results[key] = append(checker.Results[key], value...)
	}
}

func (checker *ConfigChecker) sanitizeKey(key string) string {
	/* HTML-escaped single quote to single quote */
	key = strings.Replace(key, "&#39;", "'", -1)

	/* Double quote to single quote */
	key = strings.Replace(key, "\"", "'", -1)

	/* Strip out the cruft */
	key = strings.Replace(key, "AppConfig.", "", -1)
	key = strings.Replace(key, "ENV['", "", -1)
	key = strings.Replace(key, "']", "", -1)
	key = strings.Replace(key, "Figaro.", "", -1)

	return key
}

func (checker *ConfigChecker) sanitizePaths(paths []string) []string {
	result := []string{}

	for _, path := range paths {
		result = append(result, strings.Replace(path, checker.Path, "", -1))
	}

	return result
}
