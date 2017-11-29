package modules

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/stretchr/powerwalk"
)

/* -------------------- -------------------- */

type RailsConfigChecker struct {
	ConfigFiles []*ConfigFile
	ConfigPaths []string
	RailsPath   *string
}

func NewRailsConfigChecker(path *string) *RailsConfigChecker {
	checker := RailsConfigChecker{
		RailsPath: path,
	}

	baseConfig := strings.Join([]string{*path, "config", "settings.yml"}, "/")
	checker.ConfigPaths = append(checker.ConfigPaths, baseConfig)

	return &checker
}

/* -------------------- Public Functions -------------------- */

func (railsConf *RailsConfigChecker) Keys() []string {
	exists := map[string]bool{}
	keys := []string{}

	for _, configFile := range railsConf.ConfigFiles {
		for key, _ := range configFile.Entries {
			if exists[key] != true {
				exists[key] = true
				keys = append(keys, key)
			}
		}
	}

	sort.Strings(keys)

	return keys
}

func (checker *RailsConfigChecker) Load() {
	checker.loadConfigPaths()
	checker.parseConfigFiles()
	checker.analyzeBaseConfig()
	checker.analyzeProductionConfig()

	fmt.Printf("Found %d files\n", checker.Len())
}

func (checker *RailsConfigChecker) Len() int {
	return len(checker.ConfigPaths)
}

/* -------------------- Private Functions -------------------- */

/*
* Base config (ie: settings.yml) is analyzed for the following:
* - keys that are missing values
*   We assume that the default config should not have empty values in it
 */
func (checker *RailsConfigChecker) analyzeBaseConfig() {
	baseConfig := checker.configFileByName("settings.yml")

	for _, configEntry := range baseConfig.Entries {
		configEntry.BaseIsEmpty = (configEntry.Value == "")
	}
}

/*
* Production is analyzed for the following:
* - if a hard-coded value equals a hard-coded value in any other file, warn about that
*   We assume that production should either inherit intelligently, or have unique values
 */
func (checker *RailsConfigChecker) analyzeProductionConfig() {
	prodConfig := checker.configFileByName("production.yml")

	for _, otherConfig := range checker.ConfigFiles {
		if prodConfig == otherConfig {
			continue
		}

		for _, key := range checker.Keys() {
			prodEntry := prodConfig.Entries[key]
			confEntry := otherConfig.Entries[key]

			if prodEntry == nil || confEntry == nil {
				continue
			}

			if prodEntry.Value == confEntry.Value {
				prodConfig.Entries[key].EqualsOther = true
				otherConfig.Entries[key].EqualsOther = true
			}
		}
	}
}

func (checker *RailsConfigChecker) configFileByName(name string) *ConfigFile {
	for _, configFile := range checker.ConfigFiles {
		if configFile.Name == name {
			return configFile
		}
	}
	return nil
}

func (checker *RailsConfigChecker) isYamlFile(path string) bool {
	yamlExtensions := []string{".yml", ".yaml"}
	return contains(yamlExtensions, filepath.Ext(path))
}

func (checker *RailsConfigChecker) loadConfigPaths() {
	configPath := strings.Join([]string{*checker.RailsPath, "config", "settings/"}, "/")

	var lock sync.Mutex

	powerwalk.Walk(configPath, func(path string, f os.FileInfo, err error) error {
		if checker.isYamlFile(path) {
			lock.Lock()
			defer lock.Unlock()

			checker.ConfigPaths = append(checker.ConfigPaths, path)
		}

		return nil
	})
}

/* TODO: Parallelize this operation as well */
func (checker *RailsConfigChecker) parseConfigFiles() {
	for _, path := range checker.ConfigPaths {
		fmt.Println(path)

		configFile := NewConfigFile(&path)
		if configFile.IsEmpty() == false {
			checker.ConfigFiles = append(checker.ConfigFiles, configFile)
		}
	}
}
