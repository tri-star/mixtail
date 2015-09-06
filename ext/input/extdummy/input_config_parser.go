package extdummy
import (
	"github.com/tri-star/mixtail/config"
)

const (
	INPUT_CONFIG_TYPE = "dummy"
)

type InputConfig struct {
	*config.InputBase
}

// Returns new InputSsh.
func NewInputConfig() *InputConfig {
	b := new(config.InputBase)
	i := new(InputConfig)
	i.InputBase = b
	return i
}


type InputConfigParser struct {
	*config.BaseInputConfigParser
}

func NewInputConfigParser() (ic *InputConfigParser) {
	ic = new(InputConfigParser)
	ic.Name = "dummy-input-config-parser"
	return ic
}

func (ic *InputConfigParser) CreateInputConfigFromData(name string, data map[interface{}]interface{}) (entries []config.Input, err error) {
	entry := NewInputConfig()
	entry.Name = name
	entries = append(entries, entry)
	return
}
