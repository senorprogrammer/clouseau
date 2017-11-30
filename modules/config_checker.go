package modules

import (
	"bufio"
	"fmt"
	"github.com/stretchr/powerwalk"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type ConfigChecker struct {
	Results   map[string][]string
	Path      string
	SearchStr string
}

func NewConfigChecker(path string, searchStr string) *ConfigChecker {
	checker := ConfigChecker{
		Results:   make(map[string][]string),
		Path:      path,
		SearchStr: searchStr,
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

func (checker *ConfigChecker) Run() {
	var lock sync.Mutex

	powerwalk.Walk(checker.Path, func(path string, info os.FileInfo, err error) error {
		lock.Lock()
		defer lock.Unlock()

		results := checker.find(path)
		checker.merge(results)

		return nil
	})

	fmt.Printf("Found %d config entries\n", checker.Len())
}

/* -------------------- Private Functions -------------------- */

func (checker *ConfigChecker) find(path string) *map[string][]string {
	results := make(map[string][]string)
	reg := regexp.MustCompile(checker.SearchStr)

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := reg.FindAllString(scanner.Text(), -1)

		for _, match := range matches {
			results[match] = append(results[match], path)
		}
	}

	return &results
}

func (checker *ConfigChecker) merge(results *map[string][]string) {
	if len(*results) == 0 {
		return
	}

	for key, value := range *results {
		key := checker.sanitize(key)
		value := checker.sanitizePaths(value)
		value = removeDuplicates(value)
		checker.Results[key] = append(checker.Results[key], value...)
	}
}

func (checker *ConfigChecker) sanitize(key string) string {
	/* HTML-escaped single quote to single quote */
	key = strings.Replace(key, "&#39;", "'", -1)

	/* Double quote to single quote */
	key = strings.Replace(key, "\"", "'", -1)

	/* Strip out the cruft */
	key = strings.Replace(key, "AppConfig.", "", -1)
	key = strings.Replace(key, "ENV['", "", -1)
	key = strings.Replace(key, "']", "", -1)

	return key
}

func (checker *ConfigChecker) sanitizePaths(paths []string) []string {
	result := []string{}

	/*
	* Remove the Rails path from the absolute file paths to make them shorter for display
	 */
	for _, path := range paths {
		result = append(result, strings.Replace(path, checker.Path, "", -1))
	}

	return result
}
