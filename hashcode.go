package piscine

func HashCode(dec string) string {
	size := len(dec)
	var hashed string
	for _, c := range dec {
		hash := (int(c) + size) % 127
		if hash < 32 {
			hash += 33
		}
		hashed += string(hash)
	}
	return hashed
}
