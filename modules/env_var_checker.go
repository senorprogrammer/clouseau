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

type EnvVarChecker struct {
	EnvVars   map[string][]string
	RailsPath *string
}

func NewEnvVarChecker(path *string) *EnvVarChecker {
	checker := EnvVarChecker{
		EnvVars:   make(map[string][]string),
		RailsPath: path,
	}

	return &checker
}

/* -------------------- Public Functions -------------------- */

func (checker *EnvVarChecker) Keys() []string {
	keys := make([]string, 0, len(checker.EnvVars))
	for key := range checker.EnvVars {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (checker *EnvVarChecker) Len() int {
	return len(checker.EnvVars)
}

func (checker *EnvVarChecker) Load() {
	checker.scan()

	fmt.Printf("Found %d env vars", checker.Len())
}

func (checker *EnvVarChecker) Merge(envVars *map[string][]string) {
	if len(*envVars) == 0 {
		return
	}

	for key, value := range *envVars {
		key := checker.sanitizeKey(key)
		value := checker.sanitizePaths(value)
		checker.EnvVars[key] = append(checker.EnvVars[key], value...)
	}
}

/* -------------------- Private Functions -------------------- */

/*
* Checks the file for all env var references and records them
 */
func (checker *EnvVarChecker) extractEnvVars(path string) *map[string][]string {
	envVars := make(map[string][]string)
	reg := regexp.MustCompile(`ENV\[(.*?)\]`)

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := reg.FindAllString(scanner.Text(), -1)

		for _, match := range matches {
			envVars[match] = append(envVars[match], path)
		}
	}

	return &envVars
}

func (checker *EnvVarChecker) sanitizeKey(key string) string {
	/* HTML-escaped single quote to single quote */
	key = strings.Replace(key, "&#39;", "'", -1)

	/* Double quote to single quote */
	key = strings.Replace(key, "\"", "'", -1)

	return key
}

func (checker *EnvVarChecker) sanitizePaths(paths []string) []string {
	result := []string{}

	/*
	* Remove the Rails path from the absolute file paths to make them shorter for display
	 */
	for _, path := range paths {
		result = append(result, strings.Replace(path, *checker.RailsPath, "", -1))
	}

	return result
}

func (checker *EnvVarChecker) scan() {
	var lock sync.Mutex

	powerwalk.Walk(*checker.RailsPath, func(path string, f os.FileInfo, err error) error {
		lock.Lock()
		defer lock.Unlock()

		envVars := checker.extractEnvVars(path)
		checker.Merge(envVars)

		return nil
	})
}
