package dummy
import (
	"github.com/tri-star/mixtail/config"
	"github.com/tri-star/mixtail/ext"
)


type InputConfigExtension struct {
	name string
}

func NewInputConfigExtension() (ic *InputConfigExtension) {
	ic = new(InputConfigExtension)
	ic.name = ext.INPUT_CONFIG_DUMMY
	return ic
}

func (ic *InputConfigExtension) Name() string {
	return ic.Name
}

func (ic *InputConfigExtension) CreateInputConfigFromData(name string, data map[interface{}]interface{}) (entries []config.Input, err error) {
	entry := config.NewInputSsh()
	entry.Name = name
	entries = append(entries, entry)
	return
}
