package main

import (
	"./app"
)

func main() {

	var application *app.Application

	application = app.GetInstance()
	application.Run()
}
