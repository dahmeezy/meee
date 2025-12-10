package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	res := ""
	for n > 0 {
		res = string('0'+n%10) + res
		n /= 10
	}
	return sign + res
}
