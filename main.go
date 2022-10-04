package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

const (
	title       = "sprint-reports-wails"
	description = "Reporting all the gthings"
)

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:    title,
		Width:    1024,
		Height:   768,
		Assets:   assets,
		LogLevel: logger.DEBUG,

		OnStartup: app.startup,

		// MAC specific options
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   title,
				Message: description,
				Icon:    icon,
			},
		},

		// MAC specific options
		Linux: &linux.Options{
			Icon: icon,
		},

		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
