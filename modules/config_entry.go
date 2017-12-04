package modules

type ConfigEntry struct {
	Derived bool
	Key     string
	Value   string

	/*
	* Display flags
	* TODO: Fix this, this approach is unscalable and inflexible
	 */
	BaseIsEmpty bool // Is the config entry empty in the base configuration (ie settings.yml)?
	EqualsOther bool // Does this config entry value equal the same value in another config file?
}

func NewConfigEntry(derived bool, key, value string, baseIsEmpty, equalsOther bool) *ConfigEntry {
	configEntry := ConfigEntry{
		Derived:     derived,
		Key:         key,
		Value:       value,
		BaseIsEmpty: baseIsEmpty,
		EqualsOther: equalsOther,
	}

	return &configEntry
}
