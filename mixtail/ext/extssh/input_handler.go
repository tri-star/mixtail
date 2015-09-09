package extssh

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"fmt"
	"github.com/tri-star/mixtail/mixtail/ext"
	"github.com/tri-star/mixtail/mixtail/entity"
	"github.com/tri-star/mixtail/lib"
)


type InputHandlerFactory struct {
	*lib.BaseExtensionPoint
}

func NewInputHandlerFactory() (ihf *InputHandlerFactory){
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


type InputHandler struct{
	*ext.BaseHandler

	config *InputEntry
}

func (ih *InputHandler) GetName() string {
	return ih.config.Name
}

func (ih *InputHandler) ReadInput(ch chan entity.InputData) {
	var err error

	input := entity.NewInputData()
	input.Name = ih.config.Name
	input.State = entity.INPUT_DATA_CONTINUE

	defer func() {
		if err != nil {
			input.State = entity.INPUT_DATA_END
			ih.State = ext.INPUT_STATE_ERROR
			ch <- *input
			log.Printf("Handler exited with error: %s\n", err.Error())
		}
	}()

	session, err := ih.createSession(ih.config)
	if err != nil {
		return
	}
	ih.State = ext.INPUT_STATE_RUNNING

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
				input.State = entity.INPUT_DATA_CONTINUE
			}
			ch <- *input
		}

		ih.State = ext.INPUT_STATE_DONE
		if err != nil {
			ih.State = ext.INPUT_STATE_ERROR
		} else {
			input.State = entity.INPUT_DATA_END
		}
		input.Data = nil

		// Pass a copy of "input".
		// It is needed to avoid the "input"
		// updated unfortunately by this goroutine
		// while the main thread refers.
		ch <- *input

	}()

	//Start session.
	if err := session.Run(ih.config.Command); err != nil {
		return
	}

	ih.State = ext.INPUT_STATE_DONE
}

func (ih *InputHandler) createSession(config *InputEntry) (session *ssh.Session, err error) {

	var authMethod []ssh.AuthMethod
	var key *ssh.Signer
	if(config.Identity != "") {
		key, err = ih.parsePrivateKey(config.Identity)
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

func (ih *InputHandler) parsePrivateKey(keyPath string) (key *ssh.Signer, err error) {
	buff, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	parsedKey, err := ssh.ParsePrivateKey(buff)
	key = &parsedKey
	return
}
