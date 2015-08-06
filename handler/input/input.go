package input

const (
	STATUS_WAIT = iota
	STATUS_OK
	STATUS_ERROR
)

type InputHandler interface {
	Name() string
	Type() string
	Status() uint8
	Error() error

	ReadInput(ch chan []byte)
}

type BaseHandler struct {
	name string
	typeName string
	status uint8
	err error
}


func (b *BaseHandler) Name() string {
	return b.name
}

func (b *BaseHandler) Type() string {
	return b.typeName
}

func (b *BaseHandler) Status() uint8 {
	return b.status
}

func (b *BaseHandler) Error() uint8 {
	return b.err
}

func (b *BaseHandler) ReadInput(ch chan []byte) {
}

