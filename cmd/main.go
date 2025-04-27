package main

import (
	"strings"
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type CharPosition struct {
	Char     string
	Position int
}

func GetSpecialCharactersWithPositions(input string) []CharPosition {
	var result []CharPosition

	for i, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			result = append(result, CharPosition{Char: string(char), Position: i})
		}
	}

	return result
}

func GetSpecialCharacters(position int, chars []CharPosition) string {
	var result string

	for _, char := range chars {
		if char.Position == position {
			result += char.Char
		}
	}

	return result
}
func GetPositionNotCharactersEspecial(input string) int {
	var position int
	for i, char := range input {
		if unicode.IsLetter(char) && unicode.IsDigit(char) {
			position = i
			break
		}
	}

	return position
}

func GetLastSpecialCharacter(input string, chars []CharPosition) int {
	var lastPosition int

	for _, char := range chars {
		if input == char.Char {
			lastPosition = char.Position
		}
	}

	return lastPosition
}

func GetPositionCharactersEspecial(input string, chars []CharPosition) int {
	var position int
	for _, char := range chars {
		if input == char.Char {
			position = char.Position
			break
		}
	}

	return position
}
func main() {
	// Inicializa o aplicativo Fyne
	myApp := app.New()
	myWindow := myApp.NewWindow("Máscara de CPF")

	// Campo de entrada com máscara
	input := widget.NewEntry()
	input.SetPlaceHolder("Digite o CPF")
	SetMask(input, "###.###s", true)

	// Configura o layout da janela
	myWindow.SetContent(input)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
}

func SetMask(input *widget.Entry, mask string, onlyDigits bool) {
	input.OnChanged = func(text string) {
		var numericText string

		if onlyDigits {
			numericText = strings.Map(func(r rune) rune {
				if r >= '0' && r <= '9' {
					return r
				}
				return -1
			}, text)
		}

		if len(numericText) > len(mask) {
			numericText = numericText[:len(mask)]
		}

		// Aplicar a mácara
		var maskedText string
		numericIndex := 0

		for _, maskChar := range mask {
			if numericIndex >= len(numericText) {
				break
			}

			if maskChar == '#' { // Substitui o marcador '#' pelo próximo número de numericText
				maskedText += string(numericText[numericIndex])
				numericIndex++
			} else { // Mantém os caracteres especiais da máscara
				maskedText += string(maskChar)
			}
		}

		input.SetText(maskedText)
		input.CursorColumn = len(maskedText)
	}
}
