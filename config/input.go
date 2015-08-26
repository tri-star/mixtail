package config
import (
	"errors"
)

// Input data types.
const(
	INPUT_TYPE_DUMMY = "dummy"
	INPUT_TYPE_SSH = "ssh"
)

// Input is common interface of input config.
// It is correspond to an entry of "input" section of YAML config.
type Input interface {
	GetName() string
	GetType() string

	BuildFromData(name string, data map[interface{}]interface{}) (err error)
}

type InputBase struct {
	Name string
	Type string
}

func (i *InputBase) GetName() string {
	return i.Name
}

func (i *InputBase) GetType() string {
	return i.Type
}

// BuildFromData initializes its data with given data.
// The data contains YAML data which parsed by yaml library.
// (yaml library returns the result as type of map[interface{}]interface{})
//
// This method handles common initialization process.
// All sub classes have to call this method.
func (i *InputBase) BuildFromData(name string, data map[interface{}]interface{}) (err error) {
	i.Name = name

	var ok bool
	i.Type, ok = data["type"].(string)
	if !ok {
		err = errors.New(name + ": 'type' is not specified.")
	}

	return
}

// InputRemote implements config entry that uses ssh.
type InputRemote struct {
	*InputBase

	Host string
	Port uint16
	User string
	Pass string
	Identity string
	Command string
}

// Returns new InputRemote.
func NewInputRemote() *InputRemote{
	b := new(InputBase)
	i := new(InputRemote)
	i.InputBase = b
	return i
}

// Initialize its self by given data.
func (ir *InputRemote) BuildFromData(name string, data map[interface{}]interface{}) (err error) {
	ir.InputBase.BuildFromData(name, data)

	var ok bool
	ir.Host, ok = data["host"].(string)
	if !ok {
		err = errors.New(name + ": 'host' is not specified.")
		return
	}
	ir.User, ok = data["user"].(string)
	if !ok {
		err = errors.New(name + ": 'user' is not specified.")
		return
	}
	ir.Command, ok = data["command"].(string)
	if !ok {
		err = errors.New(name + ": 'command' is not specified.")
		return
	}

	//Optional fields
	port, ok := data["port"].(uint16)
	if ok {
		ir.Port = port
	}
	pass, ok := data["pass"].(string)
	if ok {
		ir.Pass = pass
	}
	identity, ok := data["identity"].(string)
	if ok {
		ir.Identity = identity
	}

	return
}
