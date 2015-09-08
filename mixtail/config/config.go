package config


type Config struct{

	Inputs []Input

	Logging bool
	LogPath string
}

// Returns new Config.
func NewConfig() *Config {
	config := new(Config)

	config.Inputs = make([]Input, 0)
	config.Logging = false

	return config
}
