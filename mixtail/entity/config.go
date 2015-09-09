package entity


type Config struct{

	InputEntries []InputEntry

	Logging bool
	LogPath string
}

// Returns new Config.
func NewConfig() *Config {
	config := new(Config)

	config.InputEntries = make([]InputEntry, 0)
	config.Logging = false

	return config
}
