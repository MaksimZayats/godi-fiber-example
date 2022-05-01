package main

import (
	"github.com/MaximZayats/godi-fiber-example/pkg/application"
)

func main() {
	app := application.GetApp()
	_ = app.Listen(":3000")
}
