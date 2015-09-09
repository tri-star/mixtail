package main

import (
	"github.com/tri-star/mixtail/mixtail"
	"github.com/tri-star/mixtail/lib"
	"github.com/tri-star/mixtail/mixtail/ext/extssh"
)

func main() {

	extensionManager := lib.NewExtensionManager()
	extssh.NewExtension().InstallExtensionPoints(extensionManager)

	application := mixtail.NewApplication(extensionManager)
	application.Run()
}
