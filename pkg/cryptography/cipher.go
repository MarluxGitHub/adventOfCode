package cryptography

func Cipher(s string, shift int) string {
	var result string
	for _, c := range s {
		if c == '-' {
			result += " "
		} else {
			result += string(cipherChar(c, shift))
		}
	}
	return result
}

func cipherChar(c rune, shift int) rune {
	if c == '-' {
		return ' '
	}
	return rune((int(c)-97+(shift%26+26))%26 + 97)
}
