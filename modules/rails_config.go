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

type RailsConfig struct {
	ConfigFiles []*ConfigFile
	ConfigPaths []string
	RailsPath   *string
}

func NewRailsConfig(path *string) *RailsConfig {
	railsConf := RailsConfig{
		RailsPath: path,
	}

	/*
	* Defines the base configuration file from which all others will either inherit or over-ride
	 */
	baseConfig := strings.Join([]string{*path, "config", "settings.yml"}, "/")
	railsConf.ConfigPaths = append(railsConf.ConfigPaths, baseConfig)

	return &railsConf
}

/* -------------------- Public Functions -------------------- */

func (railsConf *RailsConfig) ConfigFileByName(name string) *ConfigFile {
	for _, configFile := range railsConf.ConfigFiles {
		if configFile.Name == name {
			return configFile
		}
	}
	return nil
}

func (railsConf *RailsConfig) Keys() []string {
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

func (railsConf *RailsConfig) Load(path *string) {
	railsConf.loadConfigPaths()
	railsConf.parseConfigFiles()

	fmt.Printf("Found %d files\n", railsConf.Len())
}

func (railsConf *RailsConfig) Len() int {
	return len(railsConf.ConfigPaths)
}

/* -------------------- Private Functions -------------------- */

func (railsConf *RailsConfig) isYamlFile(path string) bool {
	yamlExtensions := []string{".yml", ".yaml"}
	return contains(yamlExtensions, filepath.Ext(path))
}

func (railsConf *RailsConfig) loadConfigPaths() {
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
func (railsConf *RailsConfig) parseConfigFiles() {
	for _, path := range railsConf.ConfigPaths {
		fmt.Println(path)

		configFile := NewConfigFile(&path)
		if configFile.IsEmpty() == false {
			railsConf.ConfigFiles = append(railsConf.ConfigFiles, configFile)
		}
	}
}
