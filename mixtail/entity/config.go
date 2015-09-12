package entity


type Config struct{

	//DefaultCredential stores default credential info
	// which used by ssh extension.
	// If credential info omitted at YAML confi,
	// credential info decided following order.
	// 1. If the host name matches one of "Hosts" map entry, "Hosts" map entry is used.
	// 2. Otherwise, DefaultCredential is used.
	// 3. If DefaultCredential is not specified, parameter error raised.
	//
	// Currently this setting is defined at application layer.
	DefaultCredential *Credential

	//Hosts stores default credential info per hosts.
	Hosts map[string]*Credential

	InputEntries []InputEntry

	Logging bool
	LogPath string
}

// Returns new Config.
func NewConfig() *Config {
	config := new(Config)

	config.DefaultCredential = new(Credential)
	config.Hosts = make(map[string]*Credential)
	config.InputEntries = make([]InputEntry, 0)
	config.Logging = false

	return config
}


func (c *Config) GetDefaultCredential(hostName string) (cred *Credential) {
	cred, found := c.Hosts[hostName]
	if !found {
		cred = c.DefaultCredential
	}
	return
}
