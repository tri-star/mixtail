package extdummy
import (
	"github.com/tri-star/mixtail/mixtail/config"
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/ext"
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
	*lib.BaseExtension
}

func NewInputConfigParser() (ic *InputConfigParser) {
	ic = new(InputConfigParser)
	ic.BaseExtension = new(lib.BaseExtension)
	ic.Name = "dummy"
	return ic
}

func (ic *InputConfigParser) InstallExtension(em *lib.ExtensionManager)  {
	em.RegisterExtension("dummy", ic)
}

func (ic *InputConfigParser) CreateInputConfigFromData(name string, data map[interface{}]interface{}) (entries []config.Input, err error) {
	entry := NewInputConfig()
	entry.Name = name
	entries = append(entries, entry)
	return
}
