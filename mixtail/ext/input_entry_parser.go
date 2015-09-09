package ext
import (
	"github.com/tri-star/mixtail/mixtail/entity"
	"github.com/tri-star/mixtail/lib"
)


type InputEntryParser interface {
	lib.ExtensionPoint
	CreateInputEntriesFromData(name string, data map[interface{}]interface{}) (entries []entity.InputEntry, err error)
}
