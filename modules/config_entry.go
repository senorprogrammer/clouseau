package modules

type ConfigEntry struct {
	Key         string
	Value       string
	EqualsOther bool
}

func NewConfigEntry(key, value string, equalsOther bool) *ConfigEntry {
	configEntry := ConfigEntry{
		Key:         key,
		Value:       value,
		EqualsOther: equalsOther,
	}

	return &configEntry
}
