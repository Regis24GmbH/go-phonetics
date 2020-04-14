package gophonetics

import (
	"regexp"
	"strings"

	godiacritics "gopkg.in/Regis24GmbH/go-diacritics.v2"
)

// NewPhoneticCode takes a word an returns the phonetic code.
func NewPhoneticCode(word string) string {
	code := ""

	// only lower case
	word = strings.ToLower(word)

	// Transform the following characters: v->f, w->f, j->i, y->i, ph->f, ä->a,
	// ö->o, ü->u, ß->ss, é->e, è->e, ê->e, à->a, á->a, â->a, ë->e
	word = replaceChars(godiacritics.Normalize(word))

	// only letters (no numbers, no special characters)
	reg := regexp.MustCompile("[^a-z]+")
	word = reg.ReplaceAllString(word, "")

	wordlen := len(word)
	if wordlen == 0 {
		return ""
	}
	char := strings.Split(word, "")

	var x int

	// special cases
	if char[0] == "c" {
		code = "8"
		if wordlen > 1 {
			switch char[1] {
			case "a", "h", "k", "l", "o", "q", "r", "u", "x":
				code = "4"
			}
		}
		x = 1
	} else {
		x = 0
	}

	for ; x < wordlen; x++ {
		switch char[x] {
		case "a", "e", "i", "o", "u":
			code += "0"
		case "b", "p":
			code += "1"
		case "d", "t":
			if x+1 < wordlen {
				switch char[x+1] {
				case "c", "s", "z":
					code += "8"
				default:
					code += "2"
				}
			} else {
				code += "2"
			}
		case "f":
			code += "3"
		case "g", "k", "q":
			code += "4"
		case "c":
			if x+1 < wordlen {
				switch char[x+1] {
				case "a", "h", "k", "o", "q", "u", "x":
					switch char[x-1] {
					case "s", "z":
						code += "8"
					default:
						code += "4"
					}
				default:
					code += "8"
				}
			} else {
				code += "8"
			}
		case "x":
			if x > 0 {
				switch char[x-1] {
				case "c", "k", "q":
					code += "8"
				default:
					code += "48"
				}
			} else {
				code += "48"
			}
		case "l":
			code += "5"
		case "m", "n":
			code += "6"
		case "r":
			code += "7"
		case "s", "z":
			code += "8"
		}
	}

	// remove all "0", except at the beginning
	codelen := len(code)
	if codelen == 0 {
		return ""
	}
	codeChars := strings.Split(code, "")
	phoneticcode := codeChars[0]

	for x := 1; x < codelen; x++ {
		if codeChars[x] != "0" {
			phoneticcode += codeChars[x]
		}
	}

	return removeDuplicates(phoneticcode)
}

//--------------------------------------------------------------------------

func removeDuplicates(oldString string) string {
	// Add \122 at the end, otherwise there is one char missing; see 'hsz':
	// http://stackoverflow.com/questions/7780794/javascript-regex-remove-duplicate-characters
	oldString += "\122"
	oldStringSlice := strings.Split(oldString, "")
	oldlen := len(oldStringSlice)
	newString := ""

	char := oldStringSlice[0]

	for i := 1; i < oldlen; i++ {
		if char != oldStringSlice[i] {
			newString += char
		}

		char = oldStringSlice[i]
	}

	return newString
}

//--------------------------------------------------------------------------

func replaceChars(word string) string {
	oldChars := []string{"v", "w", "j", "y", "ph"}
	newChars := []string{"f", "f", "i", "i", "f"}

	for idx, char := range oldChars {
		word = strings.Replace(word, char, newChars[idx], -1)
	}

	return word
}

//--------------------------------------------------------------------------
