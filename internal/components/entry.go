package components

import "fyne.io/fyne/v2/widget"

func NewEnrty(placeHolder string) *widget.Entry {
	inpuut := widget.NewEntry()
	inpuut.SetPlaceHolder(placeHolder)

	return inpuut
}
