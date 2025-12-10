package piscine

func HashCode(dec string) string {
	inc := len(dec)
	var hashed string
	for i := 0; i < len(dec); i++ {
		ch := int(dec[i])
		newch := ch + inc
		hashed += string(newch)
	}
	return hashed
}
