package main

import (
	"embed"
	"log"
	"os"
	"smoovka/models"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = init_db()
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

func init_db() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlserver.Open(os.Getenv("DSN_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Pause{})

	return db, nil
}
