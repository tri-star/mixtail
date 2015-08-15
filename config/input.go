package config

const(
	INPUT_TYPE_DUMMY = "dummy"
	INPUT_TYPE_SSH = "ssh"
)

type Input interface {
	GetName() string
	GetType() string
}


type InputBase struct {
	Name string
	Type string
}

func (i *InputBase) GetName() string {
	return i.Name
}

func (i *InputBase) GetType() string {
	return i.Type
}



type InputRemote struct {
	*InputBase

	Host string
	Port uint16
	User string
	Pass string
	Identity string
	Command string
}

func NewInputRemote() *InputRemote{
	b := new(InputBase)
	i := new(InputRemote)
	i.InputBase = b
	return i
}
