package main

import (
	"embed"
	"log"

	"smoovka/services"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = services.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the wails application
	app := NewApp()

	err = wails.Run(&options.App{
		Title:  "Smoovka DevHub",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},

		Frameless: false,

		Linux: &linux.Options{
			WindowIsTranslucent: true,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "Smoovka",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
