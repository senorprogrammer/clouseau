package modules

import (
	"fmt"
	"path/filepath"
	"strings"
)

type ConfigFile struct {
	Entries map[string]string
	Name    string
	Path    *string
}

func NewConfigFile(path *string) *ConfigFile {
	configFile := ConfigFile{
		Name:    filepath.Base(*path),
		Path:    path,
		Entries: make(map[string]string),
	}

	return &configFile
}

func (configFile *ConfigFile) Append(line string) {
	if len(line) == 0 {
		return
	}

	if string(line[0]) == "#" {
		return
	}

	parts := strings.SplitN(line, ":", 2)

	if len(parts) > 1 {
		first := stripQuotes(strings.TrimSpace(parts[0]))
		last := stripQuotes(strings.TrimSpace(parts[1]))

		fmt.Printf("%s : %s\n", first, last)

		configFile.Entries[first] = last
	}
}

func (configFile *ConfigFile) Len() int {
	return len(configFile.Entries)
}
