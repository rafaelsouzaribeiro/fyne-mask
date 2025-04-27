package mask

import (
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/login/fyne-mask/pkg/mask/types"
)

type Mask struct {
	input      *widget.Entry
	mask       string
	onlyDigits bool
}

func NewMask(input *widget.Entry, mask string, onlyDigits bool) types.IMask {
	return &Mask{
		input:      input,
		mask:       mask,
		onlyDigits: onlyDigits,
	}
}
