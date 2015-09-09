package extdummy
import (
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/ext"
)

const EXTENSION_NAME = "dummy"

type Extension struct {
	*lib.BaseExtension
}


func NewExtension() (e *Extension) {
	e = new(Extension)
	e.BaseExtension = new(lib.BaseExtension)
	e.Name = EXTENSION_NAME
	return
}

func (e *Extension) InstallExtensionPoints(em *lib.ExtensionManager) {
	em.RegisterExtensionPoint(ext.POINT_INPUT_CONFIG_PARSER, NewInputEntryParser())
	em.RegisterExtensionPoint(ext.POINT_INPUT_HANDLER_FACTORY, NewInputHandlerFactory())
}
