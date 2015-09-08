package extdummy
import (
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/ext"
)

type Extension struct {
	*lib.BaseExtension


}


func NewExtension() (e *Extension) {
	e = new(Extension)
	e.BaseExtension = new(lib.BaseExtension)
	return
}


func (e *Extension) GetInputConfigParser() (icp *ext.InputConfigParser) {
	icp = NewInputConfigParser()
	return
}

func (e *Extension) GetInputHandler() (ih *ext.InputHandler) {

	return
}

