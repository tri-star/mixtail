package config
import (
	"errors"
)

// Input is common interface of input config.
// It is correspond to an entry of "input" section of YAML config.
type Input interface {
	GetName() string
	GetType() string

	BuildFromData(data map[interface{}]interface{}) (err error)
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
func (i *InputBase) BuildFromData(data map[interface{}]interface{}) (err error) {
	var ok bool
	i.Name, ok = data["name"].(string)
	if !ok {
		err = errors.New(i.Name + ": 'name' is not specified.")
		return
	}
	i.Type, ok = data["type"].(string)
	if !ok {
		err = errors.New(i.Name + ": 'type' is not specified.")
		return
	}

	return
}
