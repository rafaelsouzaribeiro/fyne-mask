package mask

import (
	"regexp"

	"fyne.io/fyne/v2/widget"
)

func (v *Mask) Mask() *widget.Entry {
	v.input.OnChanged = func(text string) {
		var filteredText string
		filteredText = text

		if v.onlyDigits {
			numericMask := regexp.MustCompile(`[^0-9]+`)
			filteredText = numericMask.ReplaceAllString(text, "")
		}

		if !v.onlyDigits {
			AlphanumericMask := regexp.MustCompile(`[^a-zA-Z0-9]+`)
			filteredText = AlphanumericMask.ReplaceAllString(text, "")
		}

		if len(filteredText) > len(v.mask) {
			filteredText = filteredText[:len(v.mask)]
		}

		var maskedText string
		numericIndex := 0

		for _, maskChar := range v.mask {
			if numericIndex >= len(filteredText) {
				break
			}

			if maskChar == '#' {
				maskedText += string(filteredText[numericIndex])
				numericIndex++
			} else {
				maskedText += string(maskChar)
			}
		}

		v.input.SetText(maskedText)
		v.input.CursorColumn = len(maskedText)
	}

	return v.input
}
