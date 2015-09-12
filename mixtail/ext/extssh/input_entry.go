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
	Cred *entity.Credential
	Command string
}

// Returns new extssh.InputEntry.
func NewInputEntry() *InputEntry {
	i := new(InputEntry)
	i.InputEntryBase = new(entity.InputEntryBase)
	i.Cred = entity.NewCredential()
	return i
}

func (ie *InputEntry) GetName() string {
	return ie.Host + ":" + ie.Name
}


// Initialize its self by given data.
func (ie *InputEntry) BuildFromData(c *entity.Config, data map[interface{}]interface{}) (err error) {
	err = ie.InputEntryBase.BuildFromData(c, data)
	if err != nil {
		return
	}

	var ok bool
	ie.Host, ok = data["host"].(string)
	if !ok {
		err = errors.New(ie.Name + ": 'host' is not specified.")
		return
	}
	defaultCred := c.GetDefaultCredential(ie.Host)
	ie.Cred.User, ok = data["user"].(string)
	if !ok || ie.Cred.User == "" {
		ie.Cred.User = defaultCred.User
		if ie.Cred.User == "" {
			err = errors.New(ie.Name + ": 'user' is not specified.")
			return
		}
	}
	ie.Command, ok = data["command"].(string)
	if !ok {
		err = errors.New(ie.Name + ": 'command' is not specified.")
		return
	}

	//Optional fields
	port, ok := data["port"].(uint16)
	if ok {
		ie.Port = port
	}
	pass, ok := data["pass"].(string)
	if ok {
		ie.Cred.Pass = pass
	}
	identity, ok := data["identity"].(string)
	if ok {
		ie.Cred.Identity = identity
	}

	if ie.Cred.Pass == "" && ie.Cred.Identity == "" {
		ie.Cred.Pass = defaultCred.Pass
		ie.Cred.Identity = defaultCred.Identity
		if ie.Cred.Pass == "" && ie.Cred.Identity == "" {
			err = errors.New("password or identity must be specified.")
		}
	}
	return
}
