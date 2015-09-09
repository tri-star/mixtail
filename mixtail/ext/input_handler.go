package ext
import (
	"github.com/tri-star/mixtail/mixtail/entity"
	"github.com/tri-star/mixtail/lib"
)


const (
	INPUT_STATE_NOT_STARTED = iota
	INPUT_STATE_RUNNING
	INPUT_STATE_DONE
	INPUT_STATE_ERROR
)


type InputHandlerFactory interface {
	lib.ExtensionPoint
	NewInputHandler(entity.InputEntry) InputHandler
}


type InputHandler interface {
	GetName() string
	GetType() string
	GetState() uint8
	GetError() error

	ReadInput(ch chan entity.InputData)
}



type BaseHandler struct {
	TypeName string
	State uint8
	Err error
}


func (b *BaseHandler) GetName() string {
	return ""
}

func (b *BaseHandler) GetType() string {
	return b.TypeName
}

func (b *BaseHandler) GetState() uint8 {
	return b.State
}

func (b *BaseHandler) GetError() error {
	return b.Err
}

func (b *BaseHandler) ReadInput(ch chan *entity.InputData) {
}
