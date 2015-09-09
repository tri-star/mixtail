package entity


// Input data state.
// Which indicates whether the data is ended or not.
const (
	INPUT_DATA_CONTINUE = iota
	INPUT_DATA_END
)

// InputData is used for communicate between
// main thread and input handler's goroutine.
type InputData struct {
	Name string
	State uint8
	Data []byte
}

func NewInputData() *InputData {
	i := new(InputData)
	return i
}
