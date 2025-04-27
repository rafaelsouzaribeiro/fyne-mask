package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/internal/infra/di"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Máscara de Celular")

	content := di.NewDi()
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
}
