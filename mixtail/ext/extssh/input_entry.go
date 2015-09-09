package extssh

import (
	"errors"
	"github.com/tri-star/mixtail/mixtail/entity"
)


// extssh.InputEntry implements config entry that uses ssh.
type InputEntry struct {
	*entity.InputEntryBase

	Host string
	Port uint16
	User string
	Pass string
	Identity string
	Command string
}

// Returns new extssh.InputEntry.
func NewInputEntry() *InputEntry {
	i := new(InputEntry)
	i.InputEntryBase = new(entity.InputEntryBase)
	return i
}

func (ie *InputEntry) GetName() string {
	return ie.Host + ":" + ie.Name
}


// Initialize its self by given data.
func (ic *InputEntry) BuildFromData(data map[interface{}]interface{}) (err error) {
	err = ic.InputEntryBase.BuildFromData(data)
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
