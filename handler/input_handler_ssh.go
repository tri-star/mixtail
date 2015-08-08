package handler

import (
	"time"
	"github.com/tri-star/mixtail/config"
)

type SshHandler struct{
	*BaseHandler

	host string
	user string
	pass string
	identity string

	command string
}

func NewSshHandler(c *config.InputRemote) *SshHandler {
	b := new(BaseHandler)
	s := new(SshHandler)
	s.BaseHandler = b
	s.name = c.Name
	s.host = c.Host
	s.user = c.User
	s.pass = c.Pass
	s.identity = c.Identity
	s.command = c.Command
	s.state = INPUT_STATE_NOT_STARTED
	return s
}

func (s *SshHandler) ReadInput(ch chan *InputData) {
	input := NewInputData()
	input.Name = s.name
	input.State = INPUT_DATA_CONTINUE
	s.state = INPUT_STATE_RUNNING
	for i := 0; i < 10; i++ {
		input.Data = []byte("aaaa")
		ch <- input
		time.Sleep(100 * time.Millisecond)
	}

	s.state = INPUT_STATE_DONE
	input.State = INPUT_DATA_END
	input.Data = nil
	ch <- input
	return
}
