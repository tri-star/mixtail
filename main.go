package main

import (
	"github.com/tri-star/mixtail/mixtail"
)

func main() {

	var application *mixtail.Application

	application = mixtail.GetInstance()
	application.Run()
}
