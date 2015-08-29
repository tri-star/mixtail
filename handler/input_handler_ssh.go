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

	config *config.InputSsh
}

func NewSshHandler(c *config.InputSsh) *SshHandler {
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

func (s *SshHandler) ReadInput(ch chan InputData) {
	var err error

	input := NewInputData()
	input.Name = s.config.Host + ": " + s.config.Name
	input.State = INPUT_DATA_CONTINUE

	defer func() {
		if err != nil {
			input.State = INPUT_DATA_END
			s.state = INPUT_STATE_ERROR
			ch <- *input
			log.Printf("Handler exited with error: %s\n", err.Error())
		}
	}()

	session, err := s.createSession(s.config)
	if err != nil {
		return
	}
	s.state = INPUT_STATE_RUNNING

	//Set  a pipe for session's output.
	//Remote session will blocked(sleep) if pipe is full.
	r, err := session.StdoutPipe()
	if err != nil {
		return
	}

	//Data receiving needs start separately with main thread.
	//This thread blocks until first data received.
	go func() {
		defer session.Close()

		buffer := make([]byte, 1024)
		for {
			//r.Read will block(sleep) if pipe is empty.
			readBytes, err := r.Read(buffer)
			if err != nil {
				break
			} else {
				input.Data = buffer[:readBytes]
				input.State = INPUT_DATA_CONTINUE
			}
			ch <- *input
		}

		s.state = INPUT_STATE_DONE
		if err != nil {
			s.state = INPUT_STATE_ERROR
		} else {
			input.State = INPUT_DATA_END
		}
		input.Data = nil

		// Pass a copy of "input".
		// It is needed to avoid the "input"
		// updated unfortunately by this goroutine
		// while the main thread refers.
		ch <- *input

	}()

	//Start session.
	if err := session.Run(s.config.Command); err != nil {
		return
	}

	s.state = INPUT_STATE_DONE
//	input.State = INPUT_DATA_END
//	input.Data = nil
}

func (s *SshHandler) createSession(config *config.InputSsh) (session *ssh.Session, err error) {

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
