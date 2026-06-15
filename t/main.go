package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// "fmt"
// "piscine"

func main() {
	// piscine.QuadE(5, 3)
	// piscine.RepeatAlpha("abcd")
	// piscine.RepeatAlpha("KACHI")
	// piscine.DigitLen(20,10)

	// words:=[]string{"eat", "zainab","sad","good", "apple", "bully"}

	// sen:="the_boy_goes_to_learn"

	// fmt.Println(piscine.ToJadenCase("This website is for losers LOL!"))

	// a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	// b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]

	// fmt.Println(piscine.TestHtml())

	// fmt.Println(piscine.Isamstrong(153))
	// fmt.Println(piscine.Isamstrong(100))
	// fmt.Println(piscine.Isamstrong(370))
	// fmt.Println(piscine.Isamstrong(49))
	// fmt.Println(piscine.Isamstrong(407))
	// fmt.Println(piscine.Isamstrong(989))
	// fmt.Println(piscine.Isamstrong(1634))

	// fmt.Println(Comp(a, a))

	// fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
	// fmt.Println(Xbonacci([]int{1, 1, 1}, 10))
	// fmt.Println(High("man i need a taxi up to ubud"))

	// fmt.Println(averageScore(true, 80, 90, 100))
	// fmt.Println(duplicate_count("indivisibilities"))
	// fmt.Println(FindUniq([]float32{1, 1, 1, 1, 0.5}))
	// fmt.Println(wWave("hand"))
	// s := "abcdef"
	// fmt.Println(s[0:1])

	// fmt.Println(ToNato("A bc de fghi j k,"))

	// Revrot("1234567", 2)
	// var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	// var c = []string{"A", "B", "C", "D"}

	// fmt.Println(StockList(b, c))

	// fmt.Println(bubbleSort([]int{1, 4, 3, 1, 5, 0, 7, 6, 5, 4, 16}))
	// fmt.Println(solve([]string{"abode", "abc", "abe", "seddsf"}))

	// fmt.Println(Stati("01|15|59, 1|47|6, 01|17|20, 1|32|34, 2|3|17"))

	// fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))

	// FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})
	// fmt.Println(Thirt(1234567))
	// fmt.Println(MultiplicationTable(5))
	fmt.Println(FindChildren("ZAcbcBCz"))

}

func Tribonacci(signature [3]float64, n int) []float64 {
	// your code here
	if n == 0 {
		return []float64{}
	}
	if n <= 3 {
		return signature[:]
	}
	var res []float64

	res = append(res, signature[:]...)

	for i := 1; i <= n-3; i++ {
		if len(res) != n {
			p := len(res) - 3

			var sum float64
			for j := p; j < len(res); j++ {

				sum += res[j]

			}
			res = append(res, sum)
		}

	}
	return res

}

func Xbonacci(signature []int, n int) []int {
	h := len(signature)

	if h == 0 {
		return nil
	}
	if n == 0 {
		return []int{}
	}
	if h > n {
		return nil
	}

	var res []int

	for i := 1; i <= n-h; i++ {
		if len(res) != n {

			p := len(res) - h
			var sum int
			for j := p; j < len(res); j++ {

				sum += res[j]
			}
			res = append(res, sum)
		}
	}
	return res

}

func Comp(array1 []int, array2 []int) bool {
	// your code

	//   if array1==nil||array2==nil{
	//     return false
	//   }
	var same bool
	for i, n := range array1 {
		array1[i] = n * n
	}
	sort.Ints(array1)
	sort.Ints(array2)
	var array3 []int

	for j, v := range array2 {
		if array1[j] == array2[j] {
			array3 = append(array3, v)
		}
	}

	if len(array3) == len(array1) {
		same = true
	} else {
		same = false
	}
	return same
}
func High(s string) string {
	// your code here
	var res []int

	words := strings.Fields(s)
	for _, word := range words {
		var val int
		for _, c := range word {
			p := int(c) - 96
			val += p
		}
		res = append(res, val)

	}
	maxval := 0
	maxindex := 0

	for i, v := range res {
		if v > maxval {
			maxval = v
			maxindex = i
		}
	}

	return words[maxindex]
}

func averageScore(dropLowest bool, scores ...int) int {
	min := math.MaxInt
	sum := 0
	res := 0
	// minvalue := 0

	for _, v := range scores {
		sum += v
		if v < min {
			min = v
		}
	}

	if dropLowest {
		res = (sum - min) / (len(scores) - 1)
	} else {
		res = sum / len(scores)
	}
	return res
}

func duplicate_count(s1 string) int {
	//your code goes here
	s := strings.ToLower(s1)
	mp := make(map[rune]int)
	count := 0
	for _, c := range s {
		mp[c]++
	}

	for _, v := range mp {
		if v > 1 {
			count++
		}
	}

	return count //Instead of returning '1', you have to return integer 'i' as answer of solution.
}

func FindUniq(arr []float32) float32 {
	// Do the magic
	mp := make(map[float32]int)
	for _, v := range arr {
		mp[v]++
	}
	var diff float32
	for i, count := range mp {
		if count > 1 {
			diff = i
		}
	}

	return diff
}

func wave(words string) []string {
	// Your code here and happy coding!

	var res []string

	//   var wo []rune

	if words == "" {
		return []string{}
	}
	runes := []rune(words)

	for i, c := range runes {
		var dup []rune
		dup = append(dup, runes...)
		if c == ' ' {
			continue
		}

		dup[i] = unicode.ToUpper(dup[i])

		res = append(res, string(dup))
	}
	return res
}

func wWave(s string) []string {
	res := make([]string, len(s))
	for i := range s {
		if i > 0 {
			word := s[0:i] + strings.ToUpper(s[i:i+1]) + s[i+1:]
			res[i] = word
		} else {
			word := strings.ToUpper(s[i:i+1]) + s[i+1:]
			res[i] = word
		}
	}
	return res
}

func ToNato(words string) string {
	// the NATO map[string]string is preloaded for you
	// NATO["A"] == "Alfa", etc

	NATO := make(map[string]string)

	NATO["A"] = "Alfa"
	NATO["B"] = "Bravo"
	NATO["C"] = "Charlie"
	NATO["D"] = "Delta"
	NATO["E"] = "Echo"
	NATO["F"] = "Foxtrot"
	NATO["G"] = "Golf"
	NATO["H"] = "Hotel"
	NATO["I"] = "India"
	NATO["J"] = "Juliett"
	NATO["K"] = "Kilo"
	NATO["L"] = "lima"
	w := strings.ReplaceAll(strings.ToUpper(words), " ", "")
	var res strings.Builder
	for i, c := range w {
		res.WriteString(NATO[string(c)])
		if unicode.IsPunct(c) {
			res.WriteRune(c)
		}
		if i != len(w)-1 {
			res.WriteString(" ")
		}

	}
	return res.String()
}

// func isAnagram(s, s1 string)bool{
// 	mp:=make(map[rune]int)

// 	for
// }

func Revrot(s string, ch int) string {
	// your code

	if ch <= 0 || s == "" || ch > len(s) {
		return ""
	}

	lenstr := len(s)

	var box []string

	for i := 0; i < lenstr; i += ch {

		end := i + ch

		if end > len(s) {
			end = len(s)
		}
		chunk := s[i:end]
		if len(chunk) == ch {
			box = append(box, chunk)
		}

	}

	for i, num := range box {
		var r string
		sum := 0
		for _, d := range num {

			val, _ := strconv.Atoi(string(d))
			sum += val

		}
		if sum%2 == 0 {
			for i := ch - 1; i >= 0; i-- {
				r += string(num[i])
			}
			box[i] = r
		} else {

			for ind, d := range num {
				if ind != 0 {
					r += string(d)
				}
			}
			r += string(num[0])
			box[i] = r
		}
	}
	return strings.Join(box, "")

}

func StockList(listArt []string, listCat []string) string {
	// your code
	mp := make(map[string]int)

	var code string
	var unit int

	var res string

	for _, data := range listArt {
		part := strings.Split(data, " ")
		code = part[0]
		unit, _ = strconv.Atoi(part[1])
		for _, l := range listCat {
			mp[l] += 0
			if l == string(code[0]) {
				mp[l] += unit
			}

		}

	}

	for i, cat := range listCat {
		res += "(" + cat + " : " + strconv.Itoa(mp[cat]) + ")"
		if i < len(listCat)-1 {
			res += "-"
		}
	}
	return res

}

func bubbleSort(n []int) []int {

	for count := len(n) - 1; count > 0; count-- {
		for i, d := range n {
			temp := n[i]
			if i < len(n)-1 && d > n[i+1] {
				n[i] = n[i+1]
				n[i+1] = temp
			}
		}
	}

	return n
}

func solve(slice []string) []int {
	// Your code here and happy coding!

	var res []int

	for ind, word := range slice {
		slice[ind] = strings.ToLower(word)

		count := 0
		for i, c := range slice[ind] {
			if int(c-'a'+1) == i+1 {
				count++
			}
		}
		res = append(res, count)

	}
	return res
}

func Stati(strg string) string {
	// your code

	// if empty string is passed, return ""
	if strg == "" {
		return ""
	}
	// else split
	ts := strings.Split(strg, ", ")

	var avg string

	var median string

	var rangeInHrs string

	var val []int

	// get the actual time in seconds and store in a slice
	for _, t := range ts {
		w := strings.Split(t, "|")
		hr, _ := strconv.Atoi(w[0])
		mn, _ := strconv.Atoi(w[1])
		sec, _ := strconv.Atoi(w[2])

		totalTimeInSecs := (hr * 3600) + (mn * 60) + sec

		val = append(val, totalTimeInSecs)
	}
	valLength := len(val)

	// use bubble sort to sort slice
	for i := valLength - 1; i > 0; i-- {
		for j := range val {
			temp := val[j]
			if j < valLength-1 && val[j] > val[j+1] {
				val[j] = val[j+1]
				val[j+1] = temp
			}
		}
	}

	// get the range in int
	r := val[valLength-1] - val[0]

	// sum the values
	var sumOfVals int
	for _, v := range val {
		sumOfVals += v
	}
	// use sum to get average value
	a := sumOfVals / valLength

	// median
	md := 0
	if valLength%2 == 0 {
		v1 := val[valLength/2-1]
		v2 := val[valLength/2]
		md = (v1 + v2) / 2
	} else {
		md = val[valLength/2]
	}

	rangeInHrs = toHours(r)

	avg = toHours(a)

	median = toHours(md)

	res := "Range: " + rangeInHrs + " Average: " + avg + " Median: " + median

	return res

}

// an helper function to put the range, mean and median in required format
func toHours(val int) string {
	h := val / 3600
	m := (val % 3600) / 60
	s := (val % 3600) % 60

	hAsString := strconv.Itoa(h)
	mAsString := strconv.Itoa(m)
	sAsString := strconv.Itoa(s)

	if h >= 0 && h <= 9 {
		hAsString = "0" + hAsString
	}
	if m >= 0 && m <= 9 {
		mAsString = "0" + mAsString
	}
	if s >= 0 && s <= 9 {
		sAsString = "0" + sAsString
	}

	res := hAsString + "|" + mAsString + "|" + sAsString

	return res

}

func FindOdd(seq []int) int {
	if len(seq) == 1 {
		return seq[0] // your code here
	}

	var res int

	mp := make(map[int]int)

	for _, n := range seq {
		mp[n]++
	}

	for key, value := range mp {
		if value%2 != 0 {
			res = key
		}
	}

	return res
}

func Thirt(n int) int {
	// your code
	multipliers := []int{1, 10, 9, 12, 3, 4}

	prevVal := -1

	stationaryNum := n

	for prevVal != stationaryNum {

		prevVal = stationaryNum

		stringedNum := strconv.Itoa(stationaryNum)

		count := 0

		summedDigit := 0

		for i := len(stringedNum) - 1; i >= 0; i-- {
			digit, _ := strconv.Atoi(string(stringedNum[i]))
			multiplier := multipliers[count%6]
			summedDigit += digit * multiplier
			count++
		}
		stationaryNum = summedDigit

	}
	return stationaryNum

}

func FindChildren(dancingBrigade string) string {
	// mp:=make(map[string]string)
	// text:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// for _,c:=range text{
	// 	mp[string(c)]=strings.ToLower(string(c))
	// }

	// mp2:=make(map[string]int)

	return ""
}

