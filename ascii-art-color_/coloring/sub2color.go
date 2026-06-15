package coloring

func Sub2color(str, sub string) []bool {

	// length of the main string
	lstr := len(str)

	// how many substrings were passed (we only care if it's 0 or not).
	lsub := len(sub)

	//make a string of bool of the same length as the sentence/string
	checklist := make([]bool, lstr)

	if lsub == 0 || sub == "" {

		for i := range checklist {

			//sets all checklist to true when theres no specified substring to color
			checklist[i] = true

		}
		return checklist
	}

	//gets the substring and its length
	

	for i := 0; i < lstr; i++ {

		// checking for where part of the string matches the target and also makes sure that it doesnt overflow
		if i+len(sub) <= lstr && str[i:i+lsub] == sub {

			// make all those points that matches the target true in the checklist
			for j := i; j < i+lsub; j++ {
				checklist[j] = true
			}
		}
	}

	return checklist

}
