package translit

import (
	"strings"
	"unicode"
)

var baseRuEn = map[string]string{
	"а": "a",
	"А": "A",
	"б": "b",
	"Б": "B",
	"в": "v",
	"В": "V",
	"г": "g",
	"Г": "G",
	"д": "d",
	"Д": "D",
	"е": "e",
	"Е": "E",
	"ё": "e",
	"Ё": "E",
	"ж": "zh",
	"Ж": "Zh",
	"з": "z",
	"З": "Z",
	"и": "i",
	"И": "I",
	"й": "i",
	"Й": "I",
	"к": "k",
	"К": "K",
	"л": "l",
	"Л": "L",
	"м": "m",
	"М": "M",
	"н": "n",
	"Н": "N",
	"о": "o",
	"О": "O",
	"п": "p",
	"П": "P",
	"р": "r",
	"Р": "R",
	"с": "c",
	"С": "C",
	"т": "t",
	"Т": "T",
	"у": "u",
	"У": "U",
	"ф": "f",
	"Ф": "F",
	"х": "kh",
	"Х": "Kh",
	"ц": "ts",
	"Ц": "Ts",
	"ч": "ch",
	"Ч": "Ch",
	"ш": "sh",
	"Ш": "Sh",
	"щ": "shch",
	"Щ": "Shch",
	"ъ": "",
	"Ъ": "",
	"ы": "y",
	"Ы": "Y",
	"ь": "",
	"Ь": "",
	"э": "e",
	"Э": "E",
	"ю": "iu",
	"Ю": "Iu",
	"я": "ia",
	"Я": "Ia",
}

var baseEnRu = map[string]string{
	"a":    "а",
	"A":    "А",
	"b":    "б",
	"B":    "Б",
	"v":    "в",
	"V":    "В",
	"g":    "г",
	"G":    "Г",
	"d":    "д",
	"D":    "Д",
	"e":    "е",
	"E":    "Е",
	"h":    "х",
	"H":    "Х",
	"zh":   "ж",
	"Zh":   "Ж",
	"z":    "з",
	"Z":    "З",
	"i":    "и",
	"I":    "И",
	"k":    "к",
	"K":    "К",
	"l":    "л",
	"L":    "Л",
	"m":    "м",
	"M":    "М",
	"n":    "н",
	"N":    "Н",
	"o":    "о",
	"O":    "О",
	"p":    "п",
	"P":    "П",
	"r":    "р",
	"R":    "Р",
	"c":    "с",
	"C":    "С",
	"t":    "т",
	"T":    "Т",
	"u":    "у",
	"U":    "У",
	"f":    "ф",
	"F":    "Ф",
	"kh":   "х",
	"Kh":   "Х",
	"ts":   "ц",
	"Ts":   "Ц",
	"ch":   "ч",
	"Ch":   "Ч",
	"sh":   "ш",
	"Sh":   "Ш",
	"shch": "щ",
	"Shch": "Щ",
	"y":    "ы",
	"Y":    "Ы",
	"iu":   "ю",
	"Iu":   "Ю",
	"ia":   "я",
	"Ia":   "Я",
	"ya":   "я",
	"Ya":   "Я",
}

func isRu(rn rune) bool {
	return rn <= 'я' && rn >= 'А'
}

func TransliterateEnRu(str string) string {
	i := 0

	var res strings.Builder
	res.Grow(len(str))

	runes := []rune(str)

	for i < len(runes) {
		if i+3 < len(runes) {
			v, ok := baseEnRu[string(runes[i:i+4])]
			if ok {
				res.WriteString(v)

				i += 4
				continue
			}
		}

		if i+2 < len(runes) {
			v, ok := baseEnRu[string(runes[i:i+3])]
			if ok {
				res.WriteString(v)

				i += 3
				continue
			}
		}

		if i+1 < len(str) {
			v, ok := baseEnRu[string(runes[i:i+2])]
			if ok {
				res.WriteString(v)

				i += 2
				continue
			}
		}

		if !unicode.IsLetter(runes[i]) || isRu(runes[i]) {
			res.WriteRune(runes[i])
		} else {
			res.WriteString(baseEnRu[string(runes[i])])
		}

		i++
	}

	return res.String()
}
