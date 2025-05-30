package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/rafaelsouzaribeiro/fyne-mask/pkg/components"
	"github.com/rafaelsouzaribeiro/fyne-mask/pkg/di"
	"github.com/rafaelsouzaribeiro/fyne-mask/pkg/mask"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Enter the mask")

	content := di.NewDi(mask.NewMask(components.NewEnrty("Enter the mask"), "####-##-##", false), "Enter the mask")
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
}
