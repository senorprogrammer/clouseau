package modules

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ConfigFile struct {
	Entries map[string]*ConfigEntry
	Name    string
	Parent  *ConfigFile
	Path    *string
}

func NewConfigFile(path *string, parent *ConfigFile) *ConfigFile {
	configFile := ConfigFile{
		Entries: make(map[string]*ConfigEntry),
		Name:    filepath.Base(*path),
		Parent:  parent,
		Path:    path,
	}

	configFile.parseFile()

	return &configFile
}

/* -------------------- Public Functions -------------------- */

func (configFile *ConfigFile) Append(line string) {
	if (len(line) == 0) || (string(line[0]) == "#") {
		return
	}

	parts := strings.SplitN(line, ":", 2)

	if len(parts) > 1 {
		key := stripQuotes(strings.TrimSpace(parts[0]))
		value := stripQuotes(strings.TrimSpace(parts[1]))

		configEntry := NewConfigEntry(false, key, value, false, false)
		configFile.Entries[key] = configEntry
	}
}

func (configFile *ConfigFile) EntryAt(key string) *ConfigEntry {
	configEntry := configFile.Entries[key]

	if configEntry == nil {
		if configFile.Parent != nil && configFile != configFile.Parent {
			/* Ask the parent for the value */
			configEntry = configFile.Parent.EntryAt(key)
			configEntry.Derived = true
		} else {
			/* Else return the nil value */
			configEntry = NewConfigEntry(false, "", "", false, false)
		}
	}

	return configEntry
}

func (configFile *ConfigFile) IsEmpty() bool {
	return configFile.Len() == 0
}

func (configFile *ConfigFile) Len() int {
	return len(configFile.Entries)
}

func (configFile *ConfigFile) String() string {
	return fmt.Sprintf("%s: %d", configFile.Name, configFile.Len())
}

/* -------------------- Private Functions -------------------- */

func (configFile *ConfigFile) parseFile() {
	file, _ := os.Open(*configFile.Path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		configFile.Append(scanner.Text())
	}
}
