package cipher

func Rot13(input string) string {
	var result []rune
	for _, c := range input {
		if c >= 'A' && c <= 'Z' {
			result = append(result, 'A'+(c-'A'+13)%26)
		} else if c >= 'a' && c <= 'z' {
			result = append(result, 'a'+(c-'a'+13)%26)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

