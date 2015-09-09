package extdummy
import (
	"github.com/tri-star/mixtail/mixtail/entity"
	"github.com/tri-star/mixtail/lib"
)

type InputEntry struct {
	*entity.InputEntryBase
}


// Returns new extdummy.InputEntry.
func NewInputEntry() *InputEntry {
	i := new(InputEntry)
	i.InputEntryBase = new(entity.InputEntryBase)
	return i
}


type InputEntryParser struct {
	*lib.BaseExtensionPoint
}


func NewInputEntryParser() (iep *InputEntryParser) {
	iep = new(InputEntryParser)
	iep.BaseExtensionPoint = new(lib.BaseExtensionPoint)
	iep.Name = EXTENSION_NAME
	return
}

func (iep *InputEntryParser) CreateInputEntriesFromData(name string, data map[interface{}]interface{}) (entries []entity.InputEntry, err error) {
	entry := NewInputEntry()
	entry.Name = name
	entries = append(entries, entry)
	return
}
