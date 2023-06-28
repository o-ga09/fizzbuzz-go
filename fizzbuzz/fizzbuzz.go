package fizzbuzz

import "strconv"

func Convert(a int) string {
	var res string
	if(a % 15 == 0) {
		res = "fizzbuzz"
	} else if(a % 3 == 0) {
		res = "fizz"
	} else if(a % 5 == 0) {
		res = "buzz"
	} else {
		res = strconv.Itoa(a)
	}
	return res
}