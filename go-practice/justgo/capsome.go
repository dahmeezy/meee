package justgo

import "strings"

func Capsome(words []string, n int) []string {

	// var new []string

	// p := len(words) - N

	// for i, v := range words {

	// 	new = append(new, v)

	// 	if i >= p {

	// 		a := strings.ToUpper(v)

	// 		new[i] = a
	// 	}
	// }
	// return new

	q := make([]string, len(words))
	for i := len(words) -1; i >= 0; i-- {
		if i > n {
			q[i] = strings.ToUpper(words[i])
		}else{
			q[i] = words[i]
		}
	}
	return 	q
}
