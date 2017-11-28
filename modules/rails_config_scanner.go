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

type RailsConfigScanner struct {
	ConfigFiles []*ConfigFile
	ConfigPaths []string
	RailsPath   *string
}

func NewRailsConfigScanner(path *string) *RailsConfigScanner {
	railsConf := RailsConfigScanner{
		RailsPath: path,
	}

	baseConfig := strings.Join([]string{*path, "config", "settings.yml"}, "/")
	railsConf.ConfigPaths = append(railsConf.ConfigPaths, baseConfig)

	return &railsConf
}

/* -------------------- Public Functions -------------------- */

func (railsConf *RailsConfigScanner) Keys() []string {
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

	sort.Strings(result)

	return result
}

func (railsConf *RailsConfigScanner) Load(path *string) {
	railsConf.loadConfigPaths()
	railsConf.parseConfigFiles()
	railsConf.analyzeBaseConfig()
	railsConf.analyzeProductionConfig()

	fmt.Printf("Found %d files\n", railsConf.Len())
}

func (railsConf *RailsConfigScanner) Len() int {
	return len(railsConf.ConfigPaths)
}

/* -------------------- Private Functions -------------------- */

/*
* Base config (ie: settings.yml) is analyzed for the following:
* - keys that are missing values
*   We assume that the default config should not have empty values in it
 */
func (railsConf *RailsConfigScanner) analyzeBaseConfig() {
	baseConfig := railsConf.configFileByName("settings.yml")

	for _, configEntry := range baseConfig.Entries {
		configEntry.BaseIsEmpty = (configEntry.Value == "")
	}
}

/*
* Production is analyzed for the following:
* - if a hard-coded value equals a hard-coded value in any other file, warn about that
*   We assume that production should either inherit intelligently, or have unique values
 */
func (railsConf *RailsConfigScanner) analyzeProductionConfig() {
	prodConfig := railsConf.configFileByName("production.yml")

	for _, otherConfig := range railsConf.ConfigFiles {
		if prodConfig == otherConfig {
			continue
		}

		for _, key := range railsConf.Keys() {
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

func (railsConf *RailsConfigScanner) configFileByName(name string) *ConfigFile {
	for _, configFile := range railsConf.ConfigFiles {
		if configFile.Name == name {
			return configFile
		}
	}
	return nil
}

func (railsConf *RailsConfigScanner) isYamlFile(path string) bool {
	yamlExtensions := []string{".yml", ".yaml"}
	return contains(yamlExtensions, filepath.Ext(path))
}

func (railsConf *RailsConfigScanner) loadConfigPaths() {
	configPath := strings.Join([]string{*railsConf.RailsPath, "config", "settings/"}, "/")

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

/* TODO: Parallelize this operation as well */
func (railsConf *RailsConfigScanner) parseConfigFiles() {
	for _, path := range railsConf.ConfigPaths {
		fmt.Println(path)

		configFile := NewConfigFile(&path)
		if configFile.IsEmpty() == false {
			railsConf.ConfigFiles = append(railsConf.ConfigFiles, configFile)
		}
	}
}
