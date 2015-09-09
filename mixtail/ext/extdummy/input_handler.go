package extdummy

import (
	"time"
	"github.com/tri-star/mixtail/mixtail/ext"
	"github.com/tri-star/mixtail/mixtail/entity"
	"github.com/tri-star/mixtail/lib"
)

type InputHandlerFactory struct {
	*lib.BaseExtensionPoint
}

func NewInputHandlerFactory() (ihf *InputHandlerFactory) {
	ihf = new(InputHandlerFactory)
	ihf.BaseExtensionPoint = new(lib.BaseExtensionPoint)
	ihf.Name = EXTENSION_NAME
	return
}

func (ihf *InputHandlerFactory) NewInputHandler(ie entity.InputEntry) ext.InputHandler {
	ih := new(InputHandler)
	ih.BaseHandler = new(ext.BaseHandler)
	ih.config = ie.(*InputEntry)
	ih.State = ext.INPUT_STATE_NOT_STARTED
	return ih
}


// Dummy inplementation of InputHandler.
// This object is used for a test.
type InputHandler struct{
	*ext.BaseHandler
	config *InputEntry
}


func (ih *InputHandler) GetName() string {
	return ih.config.Name
}

func (ih *InputHandler) ReadInput(ch chan entity.InputData) {
	input := entity.NewInputData()
	input.Name = ih.GetName()
	input.State = entity.INPUT_DATA_CONTINUE
	ih.State = ext.INPUT_STATE_RUNNING
	for i := 0; i < 10; i++ {
		input.Data = []byte("aaaa")
		ch <- *input
		time.Sleep(100 * time.Millisecond)
	}

	ih.State = ext.INPUT_STATE_DONE
	input.State = entity.INPUT_DATA_END
	input.Data = nil
	ch <- *input
	return
}
