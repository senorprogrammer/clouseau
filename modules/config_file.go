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
	Path    *string
}

func NewConfigFile(path *string) *ConfigFile {
	configFile := ConfigFile{
		Entries: make(map[string]*ConfigEntry),
		Name:    filepath.Base(*path),
		Path:    path,
	}

	configFile.parseFile()

	return &configFile
}

func (configFile *ConfigFile) Append(line string) {
	/*
	* If there's an empty line, or a commented-out line, don't try to parse that line
	 */
	if (len(line) == 0) || (string(line[0]) == "#") {
		return
	}

	parts := strings.SplitN(line, ":", 2)

	if len(parts) > 1 {
		key := stripQuotes(strings.TrimSpace(parts[0]))
		value := stripQuotes(strings.TrimSpace(parts[1]))

		configEntry := NewConfigEntry(key, value, false)
		configFile.Entries[key] = configEntry
	}
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
