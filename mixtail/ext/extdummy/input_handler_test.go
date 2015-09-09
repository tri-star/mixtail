package extdummy_test

import(
	"testing"
	"github.com/tri-star/mixtail/mixtail/ext/extdummy"
	"github.com/tri-star/mixtail/mixtail/entity"
)

func TestInputHandler(t *testing.T) {

	var ic *extdummy.InputEntry
	ics := extdummy.NewInputEntry()
	ics.Name = "remote01"
	ics.Type = "dummy"

	ic = ics
	inputHandler := extdummy.NewInputHandlerFactory().NewInputHandler(ic)

	t.Logf("%s", inputHandler.GetName())

	ch := make(chan entity.InputData)
	go inputHandler.ReadInput(ch)

	inputDone := false
	for inputDone == false {
		readData := <-ch
		if(readData.State == entity.INPUT_DATA_END) {
			inputDone = true
			break
		}
		t.Logf("[%s]: %s\n", readData.Name, readData.Data)
	}
}
