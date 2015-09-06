package ext

type ExtensionManager struct {

	extensions map[string]Extension
}


type Extension interface {
	GetName() string
	GetType() string
}


func NewExtensionManager() (e *ExtensionManager) {

	e = new(ExtensionManager)
	e.Init()

	return
}


func (e *ExtensionManager) Init() {
	e.extensions = make(map[string]Extension)
}


func (e *ExtensionManager) RegisterExtension(extension Extension) {
	e.extensions[extension.GetName()] = extension
}

func (e *ExtensionManager) GetExtension(name string) (ext Extension, found bool)  {
	ext, found = e.extensions[name]
	return
}

func (e *ExtensionManager) GetExtensionsByType(typeName string) (exts []Extension)  {
	exts = make([]Extension, 0)
	for _, extension := range e.extensions {
		if(extension.GetType() == typeName) {
			exts = append(exts, extension)
		}
	}
	return
}
