package modules

type ConfigEntry struct {
	Key   string
	Value string

	/*
	* Display flags
	* TODO: Fix this, this approach is unscalable and inflexible
	 */
	BaseIsEmpty bool
	EqualsOther bool
}

func NewConfigEntry(key, value string, baseIsEmpty, equalsOther bool) *ConfigEntry {
	configEntry := ConfigEntry{
		Key:         key,
		Value:       value,
		BaseIsEmpty: baseIsEmpty,
		EqualsOther: equalsOther,
	}

	return &configEntry
}
