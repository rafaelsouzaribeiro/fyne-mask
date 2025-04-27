package types

import "fyne.io/fyne/v2/widget"

type IMask interface {
	Mask() *widget.Entry
}
