package config


type Config struct{

	Inputs []Input

}

// Returns new Config.
func NewConfig() *Config {
	config := new(Config)

	config.Inputs = make([]Input, 0)

	return config
}
