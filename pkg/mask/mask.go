package mask

import (
	"strings"

	"fyne.io/fyne/v2/widget"
)

func (v *Mask) Mask() *widget.Entry {
	v.input.OnChanged = func(text string) {
		var filteredText string

		if v.onlyDigits {
			// Permite apenas números
			filteredText = strings.Map(func(r rune) rune {
				if r >= '0' && r <= '9' {
					return r
				}
				return -1
			}, text)
		} else {
			// Permite letras e números
			filteredText = strings.Map(func(r rune) rune {
				if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
					return r
				}
				return -1
			}, text)
		}

		// Limita o texto ao tamanho da máscara
		if len(filteredText) > len(v.mask) {
			filteredText = filteredText[:len(v.mask)]
		}

		// Aplicar a máscara
		var maskedText string
		numericIndex := 0

		for _, maskChar := range v.mask {
			if numericIndex >= len(filteredText) {
				break
			}

			if maskChar == '#' { // Substitui o marcador '#' pelo próximo caractere de filteredText
				maskedText += string(filteredText[numericIndex])
				numericIndex++
			} else { // Mantém os caracteres especiais da máscara
				maskedText += string(maskChar)
			}
		}

		// Atualiza o texto no campo de entrada
		v.input.SetText(maskedText)
		v.input.CursorColumn = len(maskedText)
	}

	return v.input
}
