package config

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
	User string
	Pass string
	Identity string
	Command string
}

func NewInputRemote() *InputRemote{
	i := new(InputRemote)
	return i
}
