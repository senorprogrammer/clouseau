package modules

type ConfigEntry struct {
	Key           string
	Value         string
	EqualsDefault bool
}

func NewConfigEntry(key, value string, equalsDefault bool) *ConfigEntry {
	configEntry := ConfigEntry{
		Key:           key,
		Value:         value,
		EqualsDefault: equalsDefault,
	}

	return &configEntry
}
