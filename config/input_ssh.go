package config

import (
	"errors"
)

// InputSsh implements config entry that uses ssh.
type InputSsh struct {
	*InputBase

	Host string
	Port uint16
	User string
	Pass string
	Identity string
	Command string
}

// Returns new InputSsh.
func NewInputSsh() *InputSsh {
	b := new(InputBase)
	i := new(InputSsh)
	i.InputBase = b
	return i
}

// Initialize its self by given data.
func (is *InputSsh) BuildFromData(name string, data map[interface{}]interface{}) (err error) {
	is.InputBase.BuildFromData(name, data)

	var ok bool
	is.Host, ok = data["host"].(string)
	if !ok {
		err = errors.New(name + ": 'host' is not specified.")
		return
	}
	is.User, ok = data["user"].(string)
	if !ok {
		err = errors.New(name + ": 'user' is not specified.")
		return
	}
	is.Command, ok = data["command"].(string)
	if !ok {
		err = errors.New(name + ": 'command' is not specified.")
		return
	}

	//Optional fields
	port, ok := data["port"].(uint16)
	if ok {
		is.Port = port
	}
	pass, ok := data["pass"].(string)
	if ok {
		is.Pass = pass
	}
	identity, ok := data["identity"].(string)
	if ok {
		is.Identity = identity
	}

	return
}
