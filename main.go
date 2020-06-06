package main

import (
	"justnotes/backend"
	"os"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	var b backend.Backend

	if len(os.Args) == 2 {
		b.File = os.Args[1]
	}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "justnotes",
		JS:        js,
		CSS:       css,
		Colour:    "#ffffff",
		Resizable: true,
	})
	app.Bind(&b)
	app.Run()
}
