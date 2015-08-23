package main

import (
	"github.com/tri-star/mixtail/app"
)

func main() {

	var application *app.Application

	application = app.GetInstance()
	application.Run()
}
