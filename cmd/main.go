package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/internal/components"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/internal/infra/di"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/pkg/mask"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Enter the mask")

	content := di.NewDi(mask.NewMask(components.NewEnrty("Enter the mask"), "####-##-##", false), "Enter the mask")
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
}
