package modules

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/stretchr/powerwalk"
)

/* -------------------- -------------------- */

type RailsConfig struct {
	ConfigFiles []*ConfigFile
	ConfigPaths []string
	RailsPath   *string
}

func NewRailsConfig(path *string) *RailsConfig {
	railsConf := RailsConfig{
		RailsPath: path,
	}

	baseConfig := strings.Join([]string{*path, "config", "settings.yml"}, "/")
	railsConf.ConfigPaths = append(railsConf.ConfigPaths, baseConfig)

	return &railsConf
}

func (railsConf *RailsConfig) Check() {
	keys := railsConf.flattenKeys()
	sort.Strings(keys)

	fmt.Printf("%v", keys)
}

func (railsConf *RailsConfig) Load(path *string) {
	railsConf.loadConfigPaths()
	railsConf.parseConfigFiles()

	fmt.Printf("Found %d files\n", railsConf.Len())
}

func (railsConf *RailsConfig) Len() int {
	return len(railsConf.ConfigPaths)
}

/* -------------------- Private Functions -------------------- */

func (railsConf *RailsConfig) flattenKeys() []string {
	exists := map[string]bool{}
	result := []string{}

	for _, configFile := range railsConf.ConfigFiles {
		for key, _ := range configFile.Entries {
			if exists[key] != true {
				exists[key] = true
				result = append(result, key)
			}
		}
	}

	return result
}

func (railsConf *RailsConfig) isYamlFile(path string) bool {
	yamlExtensions := []string{".yml", ".yaml"}
	return contains(yamlExtensions, filepath.Ext(path))
}

func (railsConf *RailsConfig) loadConfigPaths() {
	configPath := strings.Join([]string{*railsConf.RailsPath, "config", "settings/"}, "/")

	fmt.Printf("Checking %s....\n", configPath)

	var lock sync.Mutex

	powerwalk.Walk(configPath, func(path string, f os.FileInfo, err error) error {
		if railsConf.isYamlFile(path) {
			lock.Lock()
			defer lock.Unlock()

			railsConf.ConfigPaths = append(railsConf.ConfigPaths, path)
		}

		return nil
	})
}

func (railsConf *RailsConfig) parseConfigFiles() {
	for _, path := range railsConf.ConfigPaths {
		file, _ := os.Open(path)
		defer file.Close()

		fmt.Println(path)

		configFile := NewConfigFile(&path)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			configFile.Append(scanner.Text())
		}

		railsConf.ConfigFiles = append(railsConf.ConfigFiles, configFile)
	}
}
