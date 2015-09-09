package lib

type ExtensionManager struct {

	extensionPoints map[string]map[string]ExtensionPoint
}


type Extension interface {
	GetName() string
	InstallExtensionPoints(em *ExtensionManager)
}

type ExtensionPoint interface {
	GetName() string
}


type BaseExtension struct {
	Name string
}

type BaseExtensionPoint struct {
	Name string
}

func NewExtensionManager() (e *ExtensionManager) {

	e = new(ExtensionManager)
	e.Init()

	return
}


func (em *ExtensionManager) Init() {
	em.extensionPoints = make(map[string]map[string]ExtensionPoint)
}


func (em *ExtensionManager) RegisterExtensionPoint(point string, extensionPoint ExtensionPoint) {
	if em.extensionPoints[point] == nil {
		em.extensionPoints[point] = make(map[string]ExtensionPoint)
	}
	name := extensionPoint.GetName()
	em.extensionPoints[point][name] = extensionPoint
}

func (em *ExtensionManager) GetExtensionPoint(point string, name string) (extensionPoint ExtensionPoint, found bool)  {
	extensionPoint, found = em.extensionPoints[point][name]
	return
}

func (em *ExtensionManager) GetExtensionPointsByType(point string) (extensionPoints map[string]ExtensionPoint, found bool)  {
	extensionPoints, found = em.extensionPoints[point]
	return
}


func (be *BaseExtension) GetName() string {
	return be.Name
}

func (be *BaseExtension) InstallExtensionPoints(em *ExtensionManager) {
}


func (bep *BaseExtensionPoint) GetName() string {
	return bep.Name
}

