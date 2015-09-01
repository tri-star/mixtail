package ext

import (
	"github.com/tri-star/mixtail/config"
	"github.com/tri-star/mixtail/handler"
)



type ExtensionManager struct {

	inputConfigs map[string]InputConfigExtension
	inputHandlers map[string]InputHandlerExtension
}


type InputConfigExtension interface {
	Name() string
	CreateInputConfigFromData(name string, data map[interface{}]interface{}) (entries []config.Input, err error)
}

type InputHandlerExtension interface {
	Name() string
	CreateInputHandler(config.Input) (handler.InputHandler, error)
}


func NewExtensionManager() (e *ExtensionManager) {

	e = new(ExtensionManager)
	e.Init()

	return
}


func (e *ExtensionManager) Init() {
	e.inputConfigs = make(map[string]InputConfigExtension)
	e.inputHandlers = make(map[string]InputHandlerExtension)
}


func (e *ExtensionManager) RegisterInputConfig(name string, input InputConfigExtension) {
	e.inputConfigs[name] = input
}

func (e *ExtensionManager) GetInputConfig(name string) (i InputConfigExtension, found bool)  {
	i, found = e.inputConfigs[name]
	return
}

func (e *ExtensionManager) GetInputConfigs() (map[string]InputConfigExtension) {
	return e.inputConfigs
}

func (e *ExtensionManager) RegisterInputHandler(name string, handler InputHandlerExtension) {
	e.inputHandlers[name] = handler
}


func (e *ExtensionManager) GetInputHandler(name string) (i InputHandlerExtension, found bool) {
	i, found = e.inputHandlers[name]
	return
}

func (e *ExtensionManager) GetInputHandlers() map[string]InputHandlerExtension {
	return e.inputHandlers
}
