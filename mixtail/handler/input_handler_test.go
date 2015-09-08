package handler_test

import(
	"testing"
	"github.com/tri-star/mixtail/mixtail/config"
	"github.com/tri-star/mixtail/mixtail/ext/input/extdummy"
	"github.com/tri-star/mixtail/mixtail/handler"
)

func TestInputHandler(t *testing.T) {

	var ic config.Input
	ics := extdummy.NewInputConfig()
	ics.Name = "remote01"
	ics.Type = extdummy.INPUT_CONFIG_TYPE

	ic = ics
	inputHandler, err := handler.NewInputHandler(ic)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Logf("%s", inputHandler.Name())

	ch := make(chan handler.InputData)
	go inputHandler.ReadInput(ch)

	inputDone := false
	for inputDone == false {
		readData := <-ch
		if(readData.State == handler.INPUT_DATA_END) {
			inputDone = true
			break
		}
		t.Logf("[%s]: %s\n", readData.Name, readData.Data)
	}
}
