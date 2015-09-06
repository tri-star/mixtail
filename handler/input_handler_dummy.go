package handler

import (
	"time"
	"github.com/tri-star/mixtail/ext/input/extdummy"
)

// Dummy inplementation of InputHandler.
// This object is used for a test.
type DummyInputHandler struct{
	*BaseHandler
	config *extdummy.InputConfig
}

func NewDummyInputHandler(c *extdummy.InputConfig) *DummyInputHandler {
	b := new(BaseHandler)
	d := new(DummyInputHandler)
	d.BaseHandler = b
	d.config = c
	d.state = INPUT_STATE_NOT_STARTED
	return d
}

func (d *DummyInputHandler) Name() string {
	return d.config.Name
}

func (d *DummyInputHandler) ReadInput(ch chan InputData) {
	input := NewInputData()
	input.Name = d.Name()
	input.State = INPUT_DATA_CONTINUE
	d.state = INPUT_STATE_RUNNING
	for i := 0; i < 10; i++ {
		input.Data = []byte("aaaa")
		ch <- *input
		time.Sleep(100 * time.Millisecond)
	}

	d.state = INPUT_STATE_DONE
	input.State = INPUT_DATA_END
	input.Data = nil
	ch <- *input
	return
}
