package handler

import (
	"errors"
	"github.com/tri-star/mixtail/config"
)

const (
	INPUT_STATE_NOT_STARTED = iota
	INPUT_STATE_RUNNING
	INPUT_STATE_DONE
	INPUT_STATE_ERROR
)

const (
	INPUT_DATA_CONTINUE = iota
	INPUT_DATA_END
)

type InputData struct {
	Name string
	State uint8
	Data []byte
}

func NewInputData() *InputData {
	i := new(InputData)
	return i
}

type InputHandler interface {
	Name() string
	Type() string
	State() uint8
	Error() error

	ReadInput(ch chan *InputData)
}

type BaseHandler struct {
	typeName string
	state uint8
	err error
}


func (b *BaseHandler) Name() string {
	return ""
}

func (b *BaseHandler) Type() string {
	return b.typeName
}

func (b *BaseHandler) State() uint8 {
	return b.state
}

func (b *BaseHandler) Error() error {
	return b.err
}

func (b *BaseHandler) ReadInput(ch chan *InputData) {
}


func NewInputHandler(c config.Input) (i InputHandler, e error) {
	i = nil
	e = nil
	switch(c.GetType()) {
	case config.INPUT_TYPE_DUMMY:
		sshConfig := c.(*config.InputRemote)
		i = NewDummyInputHandler(sshConfig)
	case config.INPUT_TYPE_SSH:
		sshConfig := c.(*config.InputRemote)
		i = NewSshHandler(sshConfig)
	default:
		e = errors.New("Unknown input type:" + c.GetType())
	}

	return
}
