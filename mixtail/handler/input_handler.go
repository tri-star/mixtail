package handler

import (
	"errors"
	"github.com/tri-star/mixtail/mixtail/config"
	"github.com/tri-star/mixtail/mixtail/ext/input/extdummy"
	"github.com/tri-star/mixtail/mixtail/ext/input/extssh"
	"github.com/tri-star/mixtail/mixtail/ext"
)



// Returns new InputHandler.
// This funtion is just factory method of InputHandler.
func NewInputHandler(c config.Input) (i ext.InputHandler, e error) {
	i = nil
	e = nil
	switch(c.GetType()) {
	case extdummy.INPUT_CONFIG_TYPE:
		config := c.(*extdummy.InputConfig)
		i = NewDummyInputHandler(config)
	case extssh.INPUT_CONFIG_TYPE:
		config := c.(*extssh.InputConfig)
		i = NewSshHandler(config)
	default:
		e = errors.New("Unknown input type:" + c.GetType())
	}

	return
}
