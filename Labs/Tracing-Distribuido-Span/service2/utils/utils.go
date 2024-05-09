package utils

import (
	"regexp"
)

func ClearSpecialCharacter(word string) string {
	re := regexp.MustCompile("[áàâãäÁÀÂÃÄéèêëÉÈÊËíìîïÍÌÎÏóòôõöÓÒÔÕÖúùûüÚÙÛÜ]")
	result := re.ReplaceAllStringFunc(word, func(s string) string {
		switch s {
		case "á", "à", "â", "ã", "ä":
			return "a"
		case "Á", "À", "Â", "Ã", "Ä":
			return "A"
		case "é", "è", "ê", "ë":
			return "e"
		case "É", "È", "Ê", "Ë":
			return "E"
		case "í", "ì", "î", "ï":
			return "i"
		case "Í", "Ì", "Î", "Ï":
			return "I"
		case "ó", "ò", "ô", "õ", "ö":
			return "o"
		case "Ó", "Ò", "Ô", "Õ", "Ö":
			return "O"
		case "ú", "ù", "û", "ü":
			return "u"
		case "Ú", "Ù", "Û", "Ü":
			return "U"
		default:
			return s
		}
	})
	return result
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}
