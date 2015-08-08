package handler

import (
	"time"
	"github.com/tri-star/mixtail/config"
)

type DummyInputHandler struct{
	*BaseHandler
}

func NewDummyInputHandler(c *config.InputRemote) *DummyInputHandler {
	b := new(BaseHandler)
	d := new(DummyInputHandler)
	d.BaseHandler = b
	d.name = c.Name
	d.state = INPUT_STATE_NOT_STARTED
	return d
}

func (d *DummyInputHandler) ReadInput(ch chan *InputData) {
	input := NewInputData()
	input.Name = d.name
	input.State = INPUT_DATA_CONTINUE
	d.state = INPUT_STATE_RUNNING
	for i := 0; i < 10; i++ {
		input.Data = []byte("aaaa")
		ch <- input
		time.Sleep(100 * time.Millisecond)
	}

	d.state = INPUT_STATE_DONE
	input.State = INPUT_DATA_END
	input.Data = nil
	ch <- input
	return
}
