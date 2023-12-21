package decrypt

func Nimbus(str string) string {
	decryptedStr := ""
	for _, c := range str {
		ascciCode := int(c)
		character := string(ascciCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
