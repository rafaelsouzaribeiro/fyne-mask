package di

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/pkg/mask/types"
)

func NewDi(mask types.IMask, label string) *fyne.Container {

	masks := mask.Mask()
	content := container.NewVBox(
		widget.NewLabel(label),
		masks,
	)

	return content
}
