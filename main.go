package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/Molsbee/crypt/assets"
	"log"
	"os"
)

var mainWindow fyne.Window

func main() {
	a := app.New()
	a.SetIcon(assets.Logo)
	setupTray(a)
	mainWindow = a.NewWindow("Crypt")
	mainWindow.SetContent(loginPage(func() {
		mainWindow.SetContent(mainApp())
	}))
	mainWindow.Resize(fyne.NewSize(640, 460))
	mainWindow.SetOnClosed(func() {
		os.Exit(0)
	})
	mainWindow.ShowAndRun()
}

func setupTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Stub", func() {})
		menu := fyne.NewMenu("", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func loginPage(onsubmit func()) *fyne.Container {
	username := widget.NewEntry()
	passphrase := widget.NewPasswordEntry()

	form := widget.NewForm(
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Passphrase", passphrase),
	)
	form.OnSubmit = onsubmit
	return container.NewPadded(form)
}

func mainApp() *container.Split {
	data := container.NewMax()

	menu := container.NewVBox(
		widget.NewButton("Images", func() {
			var images []fyne.CanvasObject
			for _, _ = range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
				images = append(images, canvas.NewImageFromResource(assets.Logo))
			}
			data.RemoveAll()
			data.Add(container.NewGridWrap(fyne.NewSize(100, 100), images...))
		}),
		widget.NewButton("Tab 2", func() {
			log.Println("Tab 2")
		}),
	)
	split := container.NewHSplit(menu, data)
	split.Offset = 0.2
	return split
}
