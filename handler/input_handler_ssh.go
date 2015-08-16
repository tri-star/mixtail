package handler

import (
	"github.com/tri-star/mixtail/config"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"fmt"
)

type SshHandler struct{
	*BaseHandler

	config *config.InputRemote
}

func NewSshHandler(c *config.InputRemote) *SshHandler {
	b := new(BaseHandler)
	s := new(SshHandler)
	s.BaseHandler = b
	s.config = c
	s.state = INPUT_STATE_NOT_STARTED
	return s
}

func (s *SshHandler) Name() string {
	return s.config.Name
}

func (s *SshHandler) ReadInput(ch chan *InputData) {
	var err error

	input := NewInputData()
	input.Name = s.config.Name
	input.State = INPUT_DATA_CONTINUE

	defer func() {
		if err != nil {
			input.State = INPUT_DATA_END
			s.state = INPUT_STATE_ERROR
			ch <- input
			log.Printf("Handler exited with error: %s\n", err.Error())
		}
	}()

	session, err := s.createSession(s.config)
	if err != nil {
		return
	}
	s.state = INPUT_STATE_RUNNING

	r, err := session.StdoutPipe()
	if err != nil {
		return
	}

	go func() {
		defer session.Close()

		buffer := make([]byte, 1024)
		for {
			readBytes, err := r.Read(buffer)
			if err != nil || readBytes == 0 {
				break
			} else {
				input.Data = buffer[:readBytes]
				input.State = INPUT_DATA_CONTINUE
			}
			ch <- input
		}

		s.state = INPUT_STATE_DONE
		if err != nil {
			s.state = INPUT_STATE_ERROR
		}
		input.State = INPUT_DATA_END
		input.Data = nil
		ch <- input
	}()
	if err := session.Run(s.config.Command); err != nil {
		return
	}

	s.state = INPUT_STATE_DONE
	input.State = INPUT_DATA_END
	input.Data = nil
}

func (s *SshHandler) createSession(config *config.InputRemote) (session *ssh.Session, err error) {

	var authMethod []ssh.AuthMethod
	var key *ssh.Signer
	if(config.Identity != "") {
		key, err = s.parsePrivateKey(config.Identity)
		if err != nil {
			return
		}
		authMethod = []ssh.AuthMethod{ssh.PublicKeys(*key),}
	} else {
		authMethod = []ssh.AuthMethod{ ssh.Password(config.Pass), }
	}

	sshConfig := new(ssh.ClientConfig)
	sshConfig.User = config.User
	sshConfig.Auth = authMethod

	port := uint16(22)
	if config.Port != 0 {
		port = config.Port
	}
	hostNameString := fmt.Sprintf("%s:%d", config.Host, port)
	client, err := ssh.Dial("tcp", hostNameString, sshConfig)
	if err != nil {
		return
	}

	session, err = client.NewSession()
	return
}

func (s *SshHandler) parsePrivateKey(keyPath string) (key *ssh.Signer, err error) {
	buff, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	parsedKey, err := ssh.ParsePrivateKey(buff)
	key = &parsedKey
	return
}
