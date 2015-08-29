package handler

import(
	"testing"
	"github.com/tri-star/mixtail/config"
)

func TestInputHandler(t *testing.T) {

	var ic config.Input
	ics := config.NewInputSsh()
	ics.Name = "remote01"
	ics.Type = config.INPUT_TYPE_DUMMY

	ic = ics
	inputHandler, err := NewInputHandler(ic)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Logf("%s", inputHandler.Name())

	ch := make(chan InputData)
	go inputHandler.ReadInput(ch)

	inputDone := false
	for inputDone == false {
		readData := <-ch
		if(readData.State == INPUT_DATA_END) {
			inputDone = true
			break
		}
		t.Logf("[%s]: %s\n", readData.Name, readData.Data)
	}
}
