package input

import (
	"../../config"
	"time"
)

type SshHandler struct{
	*InputHandler

	Config *config.InputRemote

	Command string
}

func NewSshHandler(c *config.InputRemote) *SshHandler {
	s := new(SshHandler)
	s.Config = c
	return s
}

func (s *SshHandler) ReadInput(ch chan []byte) {
	data := []byte{"aaaa"}
	for i := 0; i < 10; i++ {
		ch <- data
		time.Sleep(1 * time.Second)
	}
}
