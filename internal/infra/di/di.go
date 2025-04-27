package di

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/internal/components"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/pkg/mask"
)

func NewDi() *fyne.Container {

	cpf := mask.NewMask(components.NewEnrty("Digite o CPF"), "###.###.###-##", false)
	cpfEntry := cpf.Mask()
	content := container.NewVBox(
		widget.NewLabel("Digite o CPF:"),
		cpfEntry,
	)

	return content
}
