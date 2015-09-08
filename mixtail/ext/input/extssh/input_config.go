package extssh

import (
	"errors"
	"github.com/tri-star/mixtail/mixtail/config"
)

const (
	INPUT_CONFIG_TYPE = "ssh"
)

// ssh.InputConfig implements config entry that uses ssh.
type InputConfig struct {
	*config.InputBase

	Host string
	Port uint16
	User string
	Pass string
	Identity string
	Command string
}

// Returns new InputSsh.
func NewInputConfig() *InputConfig {
	b := new(config.InputBase)
	i := new(InputConfig)
	i.InputBase = b
	return i
}

// Initialize its self by given data.
func (ic *InputConfig) BuildFromData(data map[interface{}]interface{}) (err error) {
	err = ic.InputBase.BuildFromData(data)
	if err != nil {
		return
	}

	var ok bool
	ic.Host, ok = data["host"].(string)
	if !ok {
		err = errors.New(ic.Name + ": 'host' is not specified.")
		return
	}
	ic.User, ok = data["user"].(string)
	if !ok {
		err = errors.New(ic.Name + ": 'user' is not specified.")
		return
	}
	ic.Command, ok = data["command"].(string)
	if !ok {
		err = errors.New(ic.Name + ": 'command' is not specified.")
		return
	}

	//Optional fields
	port, ok := data["port"].(uint16)
	if ok {
		ic.Port = port
	}
	pass, ok := data["pass"].(string)
	if ok {
		ic.Pass = pass
	}
	identity, ok := data["identity"].(string)
	if ok {
		ic.Identity = identity
	}

	return
}
