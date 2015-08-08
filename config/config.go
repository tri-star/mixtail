package config


type Config struct{

	Inputs []Input

}

func NewConfig() *Config {
	config := new(Config)

	config.Inputs = make([]Input, 0)

	return config
}
