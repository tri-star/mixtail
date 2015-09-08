package lib

type ExtensionManager struct {

	extensions map[string]map[string]Extension
}


type Extension interface {
	GetName() string

	InstallExtension(em *ExtensionManager)
}


type BaseExtension struct {
	Name string
}


func NewExtensionManager() (e *ExtensionManager) {

	e = new(ExtensionManager)
	e.Init()

	return
}


func (em *ExtensionManager) Init() {
	em.extensions = make(map[string]map[string]Extension)
}


func (em *ExtensionManager) RegisterExtension(point string, extension Extension) {
	if em.extensions[point] == nil {
		em.extensions[point] = make(map[string]Extension)
	}
	name := extension.GetName()
	em.extensions[point][name] = extension
}

func (em *ExtensionManager) GetExtension(point string, name string) (extension Extension, found bool)  {
	extension, found = em.extensions[point][name]
	return
}

func (em *ExtensionManager) GetExtensionsByType(point string) (extensions map[string]Extension, found bool)  {
	return em.extensions[point]
}


func (be *BaseExtension) GetName() string {
	return be.Name
}


func (be *BaseExtension) InstallExtension(em *ExtensionManager) {
	return
}
